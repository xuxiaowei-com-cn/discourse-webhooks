package notification

import (
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

// Sender 定义通知发送接口
type Sender interface {
	// Ping 测试联通性
	Ping(discourse event.Discourse, ping interface{}, key string) error
	// SendUser 发送用户事件
	SendUser(discourse event.Discourse, user event.User, key string) error
	// SendTopic 发送话题事件
	SendTopic(discourse event.Discourse, topic event.Topic, key string) error
	// SendPost 发送帖子事件
	SendPost(discourse event.Discourse, post event.Post, key string) error
}
