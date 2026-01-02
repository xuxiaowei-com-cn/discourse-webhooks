package event

// Type EventType 定义 Discourse 事件类型的枚举
type Type string

// Discourse 事件类型常量
const (
	UserCreated        Type = "user_created"
	UserConfirmedEmail Type = "user_confirmed_email"
)

// TypeChineseMap 事件类型到中文的映射
var TypeChineseMap = map[Type]string{
	UserCreated:        "用户创建",
	UserConfirmedEmail: "用户确认邮箱",
}

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

// WeChatWorkResponse 定义企业微信 API 响应结构
type WeChatWorkResponse struct {
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

// User 结构体用于存储 Discourse Webhook 请求中的用户信息
// 该结构体将解析后的用户数据组织起来，方便日志输出
// 字段说明：
// - ID: 用户唯一标识符
// - Username: 用户名
// - Email: 用户邮箱
// - Name: 用户姓名
// - CreatedAt: 用户创建时间
// - TrustLevel: 用户信任等级
// - Moderator: 是否为版主
// - Admin: 是否为管理员
// - Locale: 用户语言设置
// - Timezone: 用户时区
// - Groups: 用户所属群组
// - UserOption: 用户选项设置
// - AvatarTemplate: 用户头像模板
// - LastPostedAt: 最后发帖时间
// - LastSeenAt: 最后在线时间
// - Muted: 是否被禁言
// - BadgeCount: 徽章数量
// - PostCount: 发帖数量
// - TimeRead: 阅读时间（秒）
// - RecentTimeRead: 最近阅读时间（秒）
// - ProfileViewCount: 个人资料查看次数
// - SecondFactorEnabled: 是否启用双因素认证
// - CanUploadProfileHeader: 是否可以上传个人资料头图
// - CanUploadUserCardBackground: 是否可以上传用户卡片背景
// - CanChatUser: 是否可以使用聊天功能
// - InvitedBy: 邀请人信息
// - PrimaryGroupId: 主要群组ID
// - PrimaryGroupName: 主要群组名称
// - FlairGroupId: 徽章群组ID
// - FlairName: 徽章名称
// - FlairUrl: 徽章URL
// - FlairBgColor: 徽章背景颜色
// - FlairColor: 徽章颜色
// - FeaturedTopic: 精选主题
// - Staged: 是否为临时用户
// - PendingCount: 待处理数量
// - MutedCategoryIds: 被禁言的分类ID列表
// - RegularCategoryIds: 普通分类ID列表
// - WatchedTags: 关注的标签列表
// - WatchingFirstPostTags: 关注首帖的标签列表
// - TrackedTags: 追踪的标签列表
// - MutedTags: 被禁言的标签列表
// - TrackedCategoryIds: 追踪的分类ID列表
// - WatchedCategoryIds: 关注的分类ID列表
// - WatchedFirstPostCategoryIds: 关注首帖的分类ID列表
// - SystemAvatarTemplate: 系统头像模板
// - MutedUsernames: 被禁言的用户名列表
// - CanMuteUsers: 是否可以禁言其他用户
// - IgnoredUsernames: 被忽略的用户名列表
// - CanIgnoreUsers: 是否可以忽略其他用户
// - AllowedPmUsernames: 允许私信的用户名列表
// - MailingListPostsPerDay: 每日邮件列表发帖数
// - UserNotificationSchedule: 用户通知计划
// - AcceptedAnswers: 被采纳的答案数量
// - FeaturedUserBadgeIds: 精选用户徽章ID列表
// - SecondaryEmails: 次要邮箱列表
// - UserNotificationSchedule: 用户通知计划
// - Groups: 用户所属群组列表
// - UserOption: 用户选项设置
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
