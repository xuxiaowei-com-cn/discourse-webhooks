package main

import (
	"bytes"
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
	"time"

	"github.com/urfave/cli/v3"
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

const (
	Name       = "discourse-webhooks"
	Usage      = "Discourse Webhook 企业微信通知服务"
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

// webhookHandler 处理 Discourse 发送的 Webhook 请求
// 参数：
// - w: http.ResponseWriter，用于向客户端返回响应
// - r: *http.Request，包含客户端发送的请求信息
// 功能：
// 1. 验证请求方法是否为 POST
// 2. 解析 URL 路径获取企业微信 Webhook key
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

	// 从 URL 路径中提取企业微信 Webhook key
	// 路径格式：/webhook/{key}
	path := r.URL.Path
	// 去除前缀 "/webhook/"，得到 key
	key := path[len("/webhook/"):]
	// 如果 key 为空，返回 400 Bad Request
	if key == "" {
		http.Error(w, "Bad request: missing webhook key", http.StatusBadRequest)
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
		log.Printf("%s: Error marshaling discourse info: %v", discourse.EventId, err)
		// 返回 500 Internal Server Error
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	// 将解析后的 Discourse 信息记录到日志
	log.Printf("%s: Discourse Info: %s", discourse.EventId, string(discourseStr))

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
	webhookSecret := os.Getenv("WEBHOOK_SECRET")
	if webhookSecret != "" {

		// 计算 HMAC-SHA256 签名
		h := hmac.New(sha256.New, []byte(webhookSecret))
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
	log.Printf("%s: Full Payload: %s", discourse.EventId, string(body))

	// 解析请求体中的用户数据
	var payload event.WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("%s: Error unmarshaling webhook payload: %v", discourse.EventId, err)
	} else {
		// 用户
		if discourse.EventType == "user" {
			// 处理用户相关事件
			switch discourse.Event {
			case "user_created", "user_confirmed_email":
				// 记录日志
				log.Printf("%s: Processing %s event - User ID: %d, Username: %s, Email: %s",
					discourse.EventId, discourse.Event, payload.User.ID, payload.User.Username, payload.User.Email)
				// 发送企业微信 Webhook 消息，传递从 URL 提取的 key
				if err := sendWeChatWorkWebhook(discourse.EventId, discourse.Event, payload.User, key); err != nil {
					log.Printf("%s: Failed to send WeChat Work webhook: %v", discourse.EventId, err)
					// 设置响应头 Content-Type 为 application/json
					w.Header().Set("Content-Type", "application/json")
					// 在响应头中添加 EventId
					w.Header().Set("X-Event-Id", discourse.EventId)
					// 返回 500 Internal Server Error 响应
					w.WriteHeader(http.StatusInternalServerError)
					// 写入 JSON 响应，包含错误信息
					w.Write([]byte(fmt.Sprintf(`%v`, err)))
					return
				}
			}
		}
	}

	// 设置响应头 Content-Type 为 application/json
	w.Header().Set("Content-Type", "application/json")
	// 在响应头中添加 EventId
	w.Header().Set("X-Event-Id", discourse.EventId)

	// 返回 200 OK 响应
	w.WriteHeader(http.StatusOK)
	// 写入 JSON 响应
	w.Write([]byte("ok"))
}

// sendWeChatWorkWebhook 发送企业微信 Webhook 消息
// 参数：
// - eventId: Webhook 事件 ID
// - eventType: 事件类型（如 user_created、user_confirmed_email 等）
// - user: 用户信息，包含 ID、用户名和邮箱
// - key: 企业微信 Webhook key，用于构建 Webhook URL
// 返回值：
// - error: 发送过程中出现的错误，如果成功则返回 nil
// 功能：
// 1. 构建 Markdown 格式的消息内容
// 2. 使用动态 key 构建企业微信 Webhook 地址
// 3. 发送 POST 请求到企业微信 Webhook 地址
// 4. 解析并验证响应
// 5. 记录发送结果日志
func sendWeChatWorkWebhook(eventId, eventType string, user event.User, key string) error {
	// 使用动态 key 构建企业微信 Webhook 地址
	webhookURL := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)

	// 获取中文事件类型
	chineseEventType, ok := event.TypeChineseMap[event.Type(eventType)]
	if !ok {
		// 如果没有找到对应的中文映射，使用英文原事件类型
		chineseEventType = eventType
	}

	// 构建 Markdown 消息内容
	message := event.WeChatWorkMessage{
		MsgType: "markdown",
	}

	message.Markdown.Content = fmt.Sprintf("### Discourse %s 事件通知\n"+
		"> **事件 ID**: %s\n"+
		"> **用户 ID**: %d\n"+
		"> **用户名**: %s\n"+
		"> **邮箱**: %s\n",
		chineseEventType, eventId, user.ID, user.Username, user.Email)

	// 将消息转换为 JSON 格式
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("%s: Error marshaling WeChat Work message: %v", eventId, err)
		return err
	}

	// 发送 POST 请求到企业微信 Webhook 地址
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Printf("%s: Error sending WeChat Work webhook: %v", eventId, err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s: Error reading WeChat Work webhook response: %v", eventId, err)
		return err
	}

	// 解析企业微信 API 响应
	var weChatResp event.WeChatWorkResponse
	if err := json.Unmarshal(respBody, &weChatResp); err != nil {
		log.Printf("%s: Error parsing WeChat Work webhook response: %v, raw response: %s", eventId, err, string(respBody))
		return err
	}

	// 验证响应状态
	if weChatResp.Errcode == 0 {
		// errcode 为 0 表示成功
		log.Printf("%s: WeChat Work webhook sent successfully, response: %s", eventId, string(respBody))
		return nil
	}

	// errcode 不为 0 表示失败
	log.Printf("%s: WeChat Work webhook sent failed, response: %s", eventId, string(respBody))
	return fmt.Errorf(string(respBody))
}

// StartCommand 启动 HTTP 服务器
func StartCommand(port string, secret string) error {
	// 设置日志格式，包含标准时间戳和文件名行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 注册 Webhook 处理函数到 "/webhook/" 路径，支持路径参数 key
	http.HandleFunc("/webhook/", func(writer http.ResponseWriter, request *http.Request) {
		webhookHandler(writer, request, secret)
	})

	// 构建服务器地址，格式为 ":端口号"
	address := fmt.Sprintf(":%s", port)
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
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "HTTP 端口",
				DefaultText: "8080",
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "secret",
				Aliases:     []string{"s"},
				Usage:       "HTTP HMAC-SHA256 签名密钥。为空代表不验证签名。",
				DefaultText: "",
				Required:    false,
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			fmt.Println("欢迎使用 Discourse Webhook 企业微信通知服务")
			var port = cmd.String("port")
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
