package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

// 定义 HTTP 客户端超时时间
const httpClientTimeout = 10 * time.Second
const WeChatAPI = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send"

// 创建带有超时设置的 HTTP 客户端
var httpClient = &http.Client{
	Timeout: httpClientTimeout,
}

type WeChatWorkSender struct{}

// FormatTime 根据系统时区格式化时间字符串
func FormatTime(args ...interface{}) template.HTML {
	// 默认时间格式
	defaultFormat := "2006-01-02 15:04:05 -07:00"

	// 检查参数数量
	if len(args) < 1 {
		return ""
	}

	// 获取时间字符串
	timeStr, ok := args[0].(string)
	if !ok {
		return ""
	}

	// 获取格式字符串，如果没有提供则使用默认格式
	format := defaultFormat
	if len(args) > 1 {
		if f, ok := args[1].(string); ok {
			format = f
		}
	}

	// 解析ISO 8601格式的时间字符串
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		// 如果解析失败，返回原始时间字符串
		return template.HTML(timeStr)
	}

	// 转换为系统本地时区
	localTime := parsedTime.Local()

	// 格式化为指定格式的时间字符串
	return template.HTML(localTime.Format(format))
}

// sendWeChatWorkRequest 发送企业微信 webhook 请求并处理响应
func sendWeChatWorkRequest(eventId, webhookURL string, message event.WeChatWorkMessage) error {
	// 将消息转换为 JSON 格式
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("%s: Error marshaling WeChat Work message: %v", eventId, err)
		return err
	}

	// 发送 POST 请求到企业微信 Webhook 地址
	resp, err := httpClient.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
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
		// errcode 为0表示成功
		log.Printf("%s: WeChat Work webhook sent successfully, response: %s", eventId, string(respBody))
		return nil
	}

	// errcode 不为0表示失败
	log.Printf("%s: WeChat Work webhook sent failed, response: %s", eventId, string(respBody))
	return fmt.Errorf("%s", string(respBody))
}

// renderTemplate 根据模板和数据渲染消息内容
func renderTemplate(templateContent string, header event.Discourse, data interface{}) (string, error) {
	// 准备模板渲染所需的数据
	tmplData := event.TemplateData{
		Header: header,
	}

	if header.EventType == "ping" {
		tmplData.Data = data.(string)
	} else {
		// 将 data 转换为 map 以便访问
		dataMap, ok := data.(map[string]interface{})
		if !ok {
			// 如果 data 不是 map，尝试转换为 JSON 再转换为 map
			jsonData, _ := json.Marshal(data)
			json.Unmarshal(jsonData, &dataMap)
		}

		tmplData.Data = dataMap
	}

	// 创建模板并注册 FormatTime 函数
	tmpl, err := template.New("webhook").Funcs(template.FuncMap{
		"FormatTime": FormatTime,
	}).Parse(templateContent)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing template: %v", err))
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, tmplData); err != nil {
		return "", errors.New(fmt.Sprintf("Error executing template: %v", err))
	}

	return buf.String(), nil
}

func (s *WeChatWorkSender) Send(header event.Discourse, data interface{}, key string) error {
	// 使用动态 key 构建企业微信 Webhook 地址
	webhookURL := fmt.Sprintf("%s?key=%s", WeChatAPI, key)

	// 构建 Markdown 消息内容
	message := event.WeChatWorkMessage{
		MsgType: "markdown",
	}

	templateContent := event.TemplateMap[event.Type(header.Event)]

	if templateContent == "" {
		return errors.New(fmt.Sprintf("the %s template string was not found", header.Event))
	}

	var err error
	// 渲染模板内容
	message.Markdown.Content, err = renderTemplate(templateContent, header, data)
	if err != nil {
		return err
	}

	// 调用公共函数发送请求
	return sendWeChatWorkRequest(header.EventId, webhookURL, message)
}
