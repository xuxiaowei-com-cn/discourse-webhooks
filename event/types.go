package event

// Type EventType 定义 Discourse 事件类型的枚举
type Type string

// Discourse 事件类型常量
const (
	AcceptedSolution         Type = "accepted_solution"
	CategoryCreated          Type = "category_created"
	CategoryUpdated          Type = "category_updated"
	GroupUpdated             Type = "group_updated"
	NotificationCreated1     Type = "notification_created1"
	NotificationCreated2     Type = "notification_created2"
	NotificationCreated4     Type = "notification_created4"
	NotificationCreated5     Type = "notification_created5"
	NotificationCreated6     Type = "notification_created6"
	NotificationCreated12    Type = "notification_created12"
	NotificationCreated14    Type = "notification_created14"
	NotificationCreated25    Type = "notification_created25"
	NotificationCreated38    Type = "notification_created38"
	Ping                     Type = "ping"
	PostCreated              Type = "post_created"
	PostDestroyed            Type = "post_destroyed"
	PostEdited               Type = "post_edited"
	PostLiked                Type = "post_liked"
	PostRecovered            Type = "post_recovered"
	ReviewableCreated        Type = "reviewable_created"
	ReviewableScoreUpdated   Type = "reviewable_score_updated"
	ReviewableTransitionedTo Type = "reviewable_transitioned_to"
	TagCreated               Type = "tag_created"
	TopicClosedStatusUpdated Type = "topic_closed_status_updated"
	TopicCreated             Type = "topic_created"
	TopicDestroyed           Type = "topic_destroyed"
	TopicEdited              Type = "topic_edited"
	TopicPinnedStatusUpdated Type = "topic_pinned_status_updated"
	TopicRecovered           Type = "topic_recovered"
	UserAddedToGroup         Type = "user_added_to_group"
	UserBadgeGranted         Type = "user_badge_granted"
	UserBadgeRevoked         Type = "user_badge_revoked"
	UserConfirmedEmail       Type = "user_confirmed_email"
	UserCreated              Type = "user_created"
	UserLoggedIn             Type = "user_logged_in"
	UserLoggedOut            Type = "user_logged_out"
	UserUpdated              Type = "user_updated"
)

// WeChatWorkMessage 结构体用于存储企业微信 Webhook 消息格式
// 该结构体将消息内容组织成企业微信 Webhook 所需的 JSON 格式
// 字段说明：
// - MsgType: 消息类型，固定为 "markdown" 用于发送 Markdown 格式消息
// - Markdown: Markdown 消息内容，包含标题和详细信息
// - At: @ 相关设置，包含 @ 所有人开关和 @ 特定用户手机号列表
type WeChatWorkMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
	At struct {
		IsAtAll   bool     `json:"isAtAll"`
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
	} `json:"at,omitempty"`
}

// WeChatWorkResponse 定义企业微信 API 响应结构
type WeChatWorkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// DingTalkMessage 结构体用于存储钉钉 Webhook 消息格式
type DingTalkMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		IsAtAll   bool     `json:"isAtAll"`
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
	} `json:"at,omitempty"`
}

// DingTalkResponse 定义钉钉 API 响应结构
type DingTalkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// Discourse 结构体用于存储 Discourse Webhook 请求的关键头信息
// 该结构体将解析后的请求头组织起来，方便后续处理
// 字段说明：
// - UserAgent: Discourse 实例的 User-Agent 头，包含 Discourse 版本信息
// - Instance: X-Discourse-Instance 头，包含 Discourse 实例的主机地址
// - EventId: X-Discourse-Event-Id 头，Webhook 事件的唯一标识符
// - EventType: X-Discourse-Event-Type 头，事件类型（如 ping、user、group_user 等）
// - Event: X-Discourse-Event 头，具体事件名称（如 ping、user_added_to_group、user_confirmed_email、user_logged_in、user_logged_out 等）
type Discourse struct {
	UserAgent string
	Instance  string
	EventId   string
	EventType string
	Event     string
}

// TemplateData 用于模板渲染的数据结构
type TemplateData struct {
	Header Discourse
	Data   interface{}
}
