package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

type WeChatWorkSender struct{}

// formatTime 根据系统时区格式化时间字符串
func formatTime(timeStr string) string {
	// 解析ISO 8601格式的时间字符串
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		// 如果解析失败，返回原始时间字符串
		return timeStr
	}

	// 转换为系统本地时区
	localTime := parsedTime.Local()

	// 格式化为易读的时间格式，包含数字时区信息
	return localTime.Format("2006-01-02 15:04:05 -07:00")
}

func (s *WeChatWorkSender) Ping(discourse event.Discourse, ping interface{}, key string) error {
	// 检查是否为测试环境，如果是则不实际发送请求
	if os.Getenv("TEST_MODE") == "true" {
		log.Printf("%s: Test mode enabled, skipping WeChat Work webhook send", discourse.EventId)
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
		discourse.EventType, discourse.EventId, ping)

	// 将消息转换为 JSON 格式
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("%s: Error marshaling WeChat Work message: %v", discourse.EventId, err)
		return err
	}

	// 发送 POST 请求到企业微信 Webhook 地址
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Printf("%s: Error sending WeChat Work webhook: %v", discourse.EventId, err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s: Error reading WeChat Work webhook response: %v", discourse.EventId, err)
		return err
	}

	// 解析企业微信 API 响应
	var weChatResp event.WeChatWorkResponse
	if err := json.Unmarshal(respBody, &weChatResp); err != nil {
		log.Printf("%s: Error parsing WeChat Work webhook response: %v, raw response: %s", discourse.EventId, err, string(respBody))
		return err
	}

	// 验证响应状态
	if weChatResp.Errcode == 0 {
		// errcode 为 0 表示成功
		log.Printf("%s: WeChat Work webhook sent successfully, response: %s", discourse.EventId, string(respBody))
		return nil
	}

	// errcode 不为 0 表示失败
	log.Printf("%s: WeChat Work webhook sent failed, response: %s", discourse.EventId, string(respBody))
	return fmt.Errorf("%s", string(respBody))
}

func (s *WeChatWorkSender) SendUser(discourse event.Discourse, user event.User, key string) error {
	// 检查是否为测试环境，如果是则不实际发送请求
	if os.Getenv("TEST_MODE") == "true" {
		log.Printf("%s: Test mode enabled, skipping WeChat Work webhook send", discourse.EventId)
		return nil
	}

	// 使用动态 key 构建企业微信 Webhook 地址
	webhookURL := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)

	// 获取中文事件类型
	chineseEventType, ok := event.TypeChineseMap[event.Type(discourse.Event)]
	if !ok {
		// 如果没有找到对应的中文映射，使用英文原事件类型
		chineseEventType = discourse.EventType
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
		chineseEventType, discourse.EventId, user.ID, user.Username, user.Email)

	// 将消息转换为 JSON 格式
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("%s: Error marshaling WeChat Work message: %v", discourse.EventId, err)
		return err
	}

	// 发送 POST 请求到企业微信 Webhook 地址
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Printf("%s: Error sending WeChat Work webhook: %v", discourse.EventId, err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s: Error reading WeChat Work webhook response: %v", discourse.EventId, err)
		return err
	}

	// 解析企业微信 API 响应
	var weChatResp event.WeChatWorkResponse
	if err := json.Unmarshal(respBody, &weChatResp); err != nil {
		log.Printf("%s: Error parsing WeChat Work webhook response: %v, raw response: %s", discourse.EventId, err, string(respBody))
		return err
	}

	// 验证响应状态
	if weChatResp.Errcode == 0 {
		// errcode 为 0 表示成功
		log.Printf("%s: WeChat Work webhook sent successfully, response: %s", discourse.EventId, string(respBody))
		return nil
	}

	// errcode 不为 0 表示失败
	log.Printf("%s: WeChat Work webhook sent failed, response: %s", discourse.EventId, string(respBody))
	return fmt.Errorf("%s", string(respBody))
}

func (s *WeChatWorkSender) SendTopic(discourse event.Discourse, topic event.Topic, key string) error {
	// 检查是否为测试环境，如果是则不实际发送请求
	if os.Getenv("TEST_MODE") == "true" {
		log.Printf("%s: Test mode enabled, skipping WeChat Work webhook send", discourse.EventId)
		return nil
	}

	// 使用动态 key 构建企业微信 Webhook 地址
	webhookURL := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)

	// 获取中文事件类型
	chineseEventType, ok := event.TypeChineseMap[event.Type(discourse.Event)]
	if !ok {
		// 如果没有找到对应的中文映射，使用英文原事件类型
		chineseEventType = discourse.EventType
	}

	// 构建话题链接
	topicLink := fmt.Sprintf("%s/t/%d", discourse.Instance, topic.ID)

	// 构建 Markdown 消息内容
	message := event.WeChatWorkMessage{
		MsgType: "markdown",
	}

	userUrl := fmt.Sprintf("%s/u/%s", discourse.Instance, topic.CreatedBy.Username)

	message.Markdown.Content = fmt.Sprintf("### Discourse %s 事件通知\n"+
		"> **事件 ID**: %s\n"+
		"> **话题 ID**: %d\n"+
		"> **话题标题**: %s\n"+
		"> **话题链接**: [%s](%s)\n"+
		"> **创建者**: [%s](%s)\n"+
		"> **创建时间**: %s\n",
		chineseEventType, discourse.EventId, topic.ID, topic.Title,
		topicLink, topicLink,
		topic.CreatedBy.Username, userUrl,
		formatTime(topic.CreatedAt))

	// 将消息转换为 JSON 格式
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("%s: Error marshaling WeChat Work message: %v", discourse.EventId, err)
		return err
	}

	// 发送 POST 请求到企业微信 Webhook 地址
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Printf("%s: Error sending WeChat Work webhook: %v", discourse.EventId, err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s: Error reading WeChat Work webhook response: %v", discourse.EventId, err)
		return err
	}

	// 解析企业微信 API 响应
	var weChatResp event.WeChatWorkResponse
	if err := json.Unmarshal(respBody, &weChatResp); err != nil {
		log.Printf("%s: Error parsing WeChat Work webhook response: %v, raw response: %s", discourse.EventId, err, string(respBody))
		return err
	}

	// 验证响应状态
	if weChatResp.Errcode == 0 {
		// errcode 为 0 表示成功
		log.Printf("%s: WeChat Work webhook sent successfully, response: %s", discourse.EventId, string(respBody))
		return nil
	}

	// errcode 不为 0 表示失败
	log.Printf("%s: WeChat Work webhook sent failed, response: %s", discourse.EventId, string(respBody))
	return fmt.Errorf("%s", string(respBody))
}

func (s *WeChatWorkSender) SendPost(discourse event.Discourse, post event.Post, key string) error {
	// 检查是否为测试环境，如果是则不实际发送请求
	if os.Getenv("TEST_MODE") == "true" {
		log.Printf("%s: Test mode enabled, skipping WeChat Work webhook send", discourse.EventId)
		return nil
	}

	// 使用动态 key 构建企业微信 Webhook 地址
	webhookURL := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)

	// 获取中文事件类型
	chineseEventType, ok := event.TypeChineseMap[event.Type(discourse.Event)]
	if !ok {
		// 如果没有找到对应的中文映射，使用英文原事件类型
		chineseEventType = discourse.EventType
	}

	// 构建 Markdown 消息内容
	message := event.WeChatWorkMessage{
		MsgType: "markdown",
	}

	// 限制帖子内容长度，避免消息过长
	content := post.Raw
	if len(content) > 200 {
		content = content[:200] + "..."
	}

	topicLink := fmt.Sprintf("%s/t/%d", discourse.Instance, post.TopicID)
	fullPostURL := topicLink + "/" + path.Base(post.PostURL)
	userUrl := fmt.Sprintf("%s/u/%s", discourse.Instance, post.Username)

	message.Markdown.Content = fmt.Sprintf("### Discourse %s 事件通知\n"+
		"> **事件 ID**: %s\n"+
		"> **话题 ID**: %d\n"+
		"> **话题标题**: %s\n"+
		"> **话题链接**: [%s](%s)\n"+
		"> **作者**: [%s](%s)\n"+
		"> **创建时间**: %s\n"+
		"> **帖子 ID**: %d\n"+
		"> **帖子内容**: %s\n"+
		"> **帖子链接**: [%s](%s)\n",
		chineseEventType,
		discourse.EventId,
		post.TopicID,
		post.TopicTitle,
		topicLink, topicLink,
		post.Username, userUrl,
		formatTime(post.CreatedAt),
		post.ID,
		content,
		fullPostURL, fullPostURL)

	// 将消息转换为 JSON 格式
	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("%s: Error marshaling WeChat Work message: %v", discourse.EventId, err)
		return err
	}

	// 发送 POST 请求到企业微信 Webhook 地址
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Printf("%s: Error sending WeChat Work webhook: %v", discourse.EventId, err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s: Error reading WeChat Work webhook response: %v", discourse.EventId, err)
		return err
	}

	// 解析企业微信 API 响应
	var weChatResp event.WeChatWorkResponse
	if err := json.Unmarshal(respBody, &weChatResp); err != nil {
		log.Printf("%s: Error parsing WeChat Work webhook response: %v, raw response: %s", discourse.EventId, err, string(respBody))
		return err
	}

	// 验证响应状态
	if weChatResp.Errcode == 0 {
		// errcode 为 0 表示成功
		log.Printf("%s: WeChat Work webhook sent successfully, response: %s", discourse.EventId, string(respBody))
		return nil
	}

	// errcode 不为 0 表示失败
	log.Printf("%s: WeChat Work webhook sent failed, response: %s", discourse.EventId, string(respBody))
	return fmt.Errorf("%s", string(respBody))
}
