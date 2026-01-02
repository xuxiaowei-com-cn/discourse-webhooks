package notification

import (
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

// Sender 定义通知发送接口
type Sender interface {
	// Ping 测试联通性
	Ping(eventId string, eventType string, ping interface{}, key string) error
	// SendUser 发送用户事件
	SendUser(eventId string, eventType string, user event.User, key string) error
}
