package notification

import (
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

type DingTalkSender struct{}

func (s *DingTalkSender) Ping(eventId string, eventType string, ping interface{}, key string) error {
	//TODO implement me
	panic("implement me")
}

func (s *DingTalkSender) SendUser(eventId string, eventType string, user event.User, key string) error {
	//TODO implement me
	panic("implement me")
}
