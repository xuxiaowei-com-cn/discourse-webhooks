package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

type WeChatWorkSender struct{}

func (s *WeChatWorkSender) Ping(eventId string, eventType string, ping interface{}, key string) error {
	// 检查是否为测试环境，如果是则不实际发送请求
	if os.Getenv("TEST_MODE") == "true" {
		log.Printf("%s: Test mode enabled, skipping WeChat Work webhook send", eventId)
		return nil
	}

	// 使用动态 key 构建企业微信 Webhook 地址
	webhookURL := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)

	// 构建 Markdown 消息内容
	message := event.WeChatWorkMessage{
		MsgType: "markdown",
	}

	message.Markdown.Content = fmt.Sprintf("### Discourse %s 事件通知\n"+
		"> **事件 ID**: %s\n"+
		"> **ping**: %v\n",
		eventType, eventId, ping)

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
	return fmt.Errorf("%s", string(respBody))
}

func (s *WeChatWorkSender) SendUser(eventId, eventType string, user event.User, key string) error {
	// 检查是否为测试环境，如果是则不实际发送请求
	if os.Getenv("TEST_MODE") == "true" {
		log.Printf("%s: Test mode enabled, skipping WeChat Work webhook send", eventId)
		return nil
	}

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
	return fmt.Errorf("%s", string(respBody))
}
