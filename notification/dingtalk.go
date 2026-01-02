package notification

import (
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

type DingTalkSender struct{}

func (d DingTalkSender) Ping(discourse event.Discourse, ping interface{}, key string) error {
	//TODO implement me
	panic("implement me")
}

func (d DingTalkSender) SendUser(discourse event.Discourse, user event.User, key string) error {
	//TODO implement me
	panic("implement me")
}

func (d DingTalkSender) SendTopic(discourse event.Discourse, topic event.Topic, key string) error {
	//TODO implement me
	panic("implement me")
}

func (d DingTalkSender) SendPost(discourse event.Discourse, post event.Post, key string) error {
	//TODO implement me
	panic("implement me")
}
