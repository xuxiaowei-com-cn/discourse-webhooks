package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
	"github.com/xuxiaowei-com-cn/discourse-webhooks/notification"
)

const (
	Name       = "discourse-webhooks"
	Usage      = "Discourse Webhook 通知服务"
	Copyright  = "徐晓伟工作室 <xuxiaowei@xuxiaowei.com.cn>"
	ProjectUrl = "github.com/xuxiaowei-com-cn/discourse-webhooks"
)

var Authors = []any{
	"徐晓伟 <xuxiaowei@xuxiaowei.com.cn>",
}

var (
	GitCommit     = ""
	GitBranch     = ""
	GitTag        = ""
	BuildTime     = ""
	CiPipelineUrl = ""
)

func getSender(platform string) (notification.Sender, error) {
	switch platform {
	case "wechat_work", "企业微信":
		return &notification.WeChatWorkSender{}, nil
	case "dingtalk", "钉钉":
		return &notification.DingTalkSender{}, nil
	default:
		return nil, fmt.Errorf("unsupported platform: %s", platform)
	}
}

// respErr 响应错误信息
func respErr(w http.ResponseWriter, discourse *event.Discourse, err error) {
	log.Printf("%s: Failed to send webhook: %v", discourse.EventId, err)
	// 设置响应头 Content-Type 为 application/json
	w.Header().Set("Content-Type", "application/json")
	// 在响应头中添加 EventId
	w.Header().Set("X-Event-Id", discourse.EventId)

	if strings.Contains(err.Error(), "not found") {
		// 返回 404 响应
		w.WriteHeader(http.StatusNotFound)
	} else {
		// 返回 500 Internal Server Error 响应
		w.WriteHeader(http.StatusInternalServerError)
	}

	// 写入 JSON 响应，包含错误信息
	_, _ = w.Write([]byte(fmt.Sprintf(`%v`, err)))
}

// webhookHandler 处理 Discourse 发送的 Webhook 请求
// 参数：
// - w: http.ResponseWriter，用于向客户端返回响应
// - r: *http.Request，包含客户端发送的请求信息
// 功能：
// 1. 验证请求方法是否为 POST
// 2. 解析 URL 路径获取平台和 Key
// 3. 解析 Discourse 特定的请求头信息
// 4. 读取并记录请求体
// 5. 解析请求体中的用户数据
// 6. 返回 200 OK 响应
func webhookHandler(w http.ResponseWriter, r *http.Request, secret string) {
	// 只允许 POST 请求，其他方法返回 405 Method Not Allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析 URL 路径
	// 格式：/webhook/{platform}/{key}
	path := r.URL.Path
	parts := strings.Split(path, "/")

	var platform string
	var key string

	if len(parts) == 4 {
		// /webhook/{platform}/{key}
		// /webhook/wechat_work/{key}
		// /webhook/dingtalk/{key}
		platform = parts[2]
		key = parts[3]
	} else {
		http.Error(w, "Bad request: invalid path", http.StatusBadRequest)
		return
	}

	if key == "" {
		http.Error(w, "Bad request: missing webhook key", http.StatusBadRequest)
		return
	}

	sender, err := getSender(platform)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad request: %v", err), http.StatusBadRequest)
		return
	}

	// 初始化 Discourse 结构体实例，用于存储解析后的请求头
	discourse := &event.Discourse{}
	// 遍历所有请求头，提取 Discourse 相关的头信息
	actualSignature := ""
	for name, values := range r.Header {
		// 确保请求头有值
		if len(values) > 0 {
			// 根据头名称将值赋值给 Discourse 结构体对应字段
			switch name {
			case "User-Agent":
				discourse.UserAgent = values[0]
			case "X-Discourse-Instance":
				discourse.Instance = values[0]
			case "X-Discourse-Event-Id":
				discourse.EventId = values[0]
			case "X-Discourse-Event-Type":
				discourse.EventType = values[0]
			case "X-Discourse-Event":
				discourse.Event = values[0]
			case "X-Discourse-Event-Signature":
				actualSignature = values[0]
			}
		}
	}

	discourseStr, err := json.Marshal(discourse)
	if err != nil {
		// 记录错误日志
		log.Printf("%s: Error marshaling header: %v", discourse.EventId, err)
		// 返回 500 Internal Server Error
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	// 将解析后的 Discourse 信息记录到日志
	log.Printf("%s: Header: %s", discourse.EventId, string(discourseStr))

	// 读取请求体内容
	body, err := io.ReadAll(r.Body)
	// 检查读取请求体是否出错
	if err != nil {
		// 记录错误日志
		log.Printf("%s: Error reading request body: %v", discourse.EventId, err)
		// 返回 500 Internal Server Error
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// 验证签名
	if secret != "" {

		// 计算 HMAC-SHA256 签名
		h := hmac.New(sha256.New, []byte(secret))
		h.Write(body)
		expectedSignature := "sha256=" + hex.EncodeToString(h.Sum(nil))

		// 比较签名
		if !hmac.Equal([]byte(expectedSignature), []byte(actualSignature)) {
			log.Printf("%s: Invalid signature. Expected: %s, Got: %s", discourse.EventId, expectedSignature, actualSignature)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Event-Id", discourse.EventId)
			w.Header().Set("X-Expected-Signature", expectedSignature)
			w.Header().Set("X-Actual-Signature", actualSignature)

			http.Error(w, "Invalid signature", http.StatusUnauthorized)
			return
		}
	}

	log.Printf("%s: Signature verified successfully", discourse.EventId)

	// 将完整的请求体记录到日志
	log.Printf("%s: Body: %s", discourse.EventId, string(body))

	// 将 body 处理为 map，根据 EventType 获取 map 中的数据并输出到日志中
	var genericPayload map[string]interface{}
	if err := json.Unmarshal(body, &genericPayload); err != nil {
		log.Printf("%s: Error unmarshaling body to map: %v", discourse.EventId, err)
		http.Error(w, "Error unmarshaling body to map", http.StatusBadRequest)
		return
	}

	data, _ := genericPayload[discourse.EventType]

	// 发送 Webhook 消息
	if err := sender.Send(*discourse, data, key); err != nil {
		respErr(w, discourse, err)
		return
	}

	// 设置响应头 Content-Type 为 application/json
	w.Header().Set("Content-Type", "application/json")
	// 在响应头中添加 EventId
	w.Header().Set("X-Event-Id", discourse.EventId)

	// 返回 200 OK 响应
	w.WriteHeader(http.StatusOK)
	// 写入 JSON 响应
	_, _ = w.Write([]byte("ok"))
}

// StartCommand 启动 HTTP 服务器
func StartCommand(port int, secret string) error {
	// 设置日志格式，包含标准时间戳和文件名行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 注册 Webhook 处理函数到 "/webhook/" 路径，支持路径参数 key
	http.HandleFunc("/webhook/", func(writer http.ResponseWriter, request *http.Request) {
		webhookHandler(writer, request, secret)
	})

	// 构建服务器地址，格式为 ":端口号"
	address := fmt.Sprintf(":%d", port)
	// 记录服务器启动日志
	log.Printf("Starting server on %s", address)
	// 启动 HTTP 服务器，监听指定地址
	// 如果服务器启动失败，则记录错误并退出程序
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	return nil
}

func main() {

	cmd := &cli.Command{
		Name:      Name,
		Usage:     Usage,
		Copyright: Copyright,
		Authors:   Authors,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "HTTP 端口",
				Value:    8080,
				Required: false,
			},
			&cli.StringFlag{
				Name:     "secret",
				Aliases:  []string{"s"},
				Usage:    "HTTP HMAC-SHA256 签名密钥。为空代表不验证签名。",
				Value:    "",
				Required: false,
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			fmt.Println("欢迎使用 Discourse Webhook 企业微信通知服务")
			var port = cmd.Int("port")
			var secret = cmd.String("secret")
			return StartCommand(port, secret)
		},
		Metadata: map[string]interface{}{
			"projectUrl": ProjectUrl,
		},
		CustomRootCommandHelpTemplate: `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} {{if .VisibleFlags}}[global options]{{end}}

{{if .Version}}VERSION:
   {{.Version}}{{end}}

{{if len .Authors}}AUTHOR:
   {{range .Authors}}{{ . }}{{end}}{{end}}

{{if .VisibleFlags}}GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
METADATA:
   ProjectUrl: {{ index .Metadata "projectUrl" }}{{ $info := call .ExtraInfo }}
   GitCommit: {{ index $info "GitCommit" }}
   GitBranch: {{ index $info "GitBranch" }}
   GitTag: {{ index $info "GitTag" }}
   BuildTime: {{ index $info "BuildTime" }}
   CiPipelineUrl: {{ index $info "CiPipelineUrl" }}
   Arch: {{ index $info "Arch" }}

COPYRIGHT:
	{{.Copyright}}
`,
		Commands: []*cli.Command{},
	}

	if GitTag != "" {
		cmd.Version = GitTag
	} else if GitCommit != "" {
		cmd.Version = GitCommit
	} else {
		cmd.Version = "dev"
	}

	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Printf("%s\n"+
			"projectUrl: %s\n"+
			"commit: %s\n"+
			"branch: %s\n"+
			"tag: %s\n"+
			"buildTime: %s\n"+
			"ciPipelineUrl: %s\n"+
			"arch: %s/%s\n",
			cmd.Version,
			ProjectUrl,
			GitCommit,
			GitBranch,
			GitTag,
			buildTimeCST(BuildTime),
			CiPipelineUrl,
			runtime.GOOS,
			runtime.GOARCH,
		)
	}

	// 使用 ExtraInfo 将元数据展示到帮助文档中
	cmd.ExtraInfo = func() map[string]string {
		return map[string]string{
			"GitCommit":     GitCommit,
			"GitBranch":     GitBranch,
			"GitTag":        GitTag,
			"BuildTime":     buildTimeCST(BuildTime),
			"CiPipelineUrl": CiPipelineUrl,
			"Project":       ProjectUrl,
			"Version":       cmd.Version,
			"Arch":          fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		}
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

// buildTimeCST 将时间转换为 CST 时区
func buildTimeCST(s string) string {
	if s == "" {
		return s
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return s
	}
	loc := time.FixedZone("CST", 8*3600)
	return t.In(loc).Format("2006-01-02T15:04:05+08:00")
}
