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

// LimitLength 限制字符串长度，超过则添加省略号
func LimitLength(args ...interface{}) template.HTML {
	// 默认最大长度
	defaultMaxLength := 1000

	// 检查参数数量
	if len(args) < 1 {
		return ""
	}

	// 获取原始字符串
	str, ok := args[0].(string)
	if !ok {
		return ""
	}

	// 获取最大长度，如果没有提供则使用默认值
	maxLength := defaultMaxLength
	if len(args) > 1 {
		if ml, ok := args[1].(int); ok {
			maxLength = ml
		}
	}

	// 如果字符串长度小于等于最大长度，直接返回
	if len(str) <= maxLength {
		return template.HTML(str)
	}

	// 截断字符串并添加省略号
	return template.HTML(str[:maxLength] + "...")
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

// nullValue 表示 Webhook 数据中的 JSON null
//
// 如果直接使用 nil 参与模板渲染，模板访问其字段（如 {{.Data.deleted_by.id}}）会返回
// “nil pointer evaluating interface {}.id” 错误，导致整条通知发送失败。
// 将 null 替换为 nullValue 后：
//   - 访问任意字段都返回空字符串，不再中断模板执行；
//   - 在 {{if}}、{{with}} 中仍然为假；
//   - 在 {{range}} 中不产生任何循环；
//   - 直接输出时通过 String 方法输出空字符串（而不是 map[]）。
type nullValue map[string]interface{}

// String 实现 fmt.Stringer，使 null 直接输出为空字符串
func (nullValue) String() string { return "" }

// normalizeNull 递归地将数据中的 nil（JSON null）替换为 nullValue
func normalizeNull(data interface{}) interface{} {
	switch value := data.(type) {
	case map[string]interface{}:
		m := make(map[string]interface{}, len(value))
		for k, v := range value {
			m[k] = normalizeNull(v)
		}
		return m
	case []interface{}:
		s := make([]interface{}, len(value))
		for i, v := range value {
			s[i] = normalizeNull(v)
		}
		return s
	case nil:
		return nullValue{}
	default:
		return value
	}
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

		tmplData.Data = normalizeNull(dataMap)
	}

	// 创建模板并注册函数
	tmpl, err := template.New("webhook").Funcs(template.FuncMap{
		"FormatTime":  FormatTime,
		"LimitLength": LimitLength,
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

// templateFor 根据事件类型获取对应的模板内容
// notification_created 事件需要再根据 notification_type 区分不同模板
func templateFor(header event.Discourse, data interface{}) string {
	if header.EventType == "notification" && header.Event == "notification_created" {
		if dataMap, ok := data.(map[string]interface{}); ok {
			notificationType := fmt.Sprintf("%v", dataMap["notification_type"])
			return event.TemplateMap[event.Type(header.Event+notificationType)]
		}
	}
	return event.TemplateMap[event.Type(header.Event)]
}

func (s *WeChatWorkSender) Send(header event.Discourse, data interface{}, key string) error {
	// 使用动态 key 构建企业微信 Webhook 地址
	webhookURL := fmt.Sprintf("%s?key=%s", WeChatAPI, key)

	// 构建 Markdown 消息内容
	message := event.WeChatWorkMessage{
		MsgType: "markdown",
	}

	templateContent := templateFor(header, data)

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
