package notification

import (
	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

type DingTalkSender struct{}

func (d DingTalkSender) Send(discourse event.Discourse, user interface{}, key string) error {
	//TODO implement me
	panic("implement me")
}
