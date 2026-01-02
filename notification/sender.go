package notification

import (
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

// Sender 定义通知发送接口
type Sender interface {
	Send(discourse event.Discourse, user interface{}, key string) error
}
