package event

var TemplateMap = map[Type]string{
	GroupUpdated: ``,
	NotificationCreated2: `# Discourse 通知创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**通知标题**: {{.Data.fancy_title}}
**通知 ID**: {{.Data.id}}
**通知类型**: {{.Data.notification_type}}
**阅读状态**: {{.Data.read}}
**高优先级**: {{.Data.high_priority}}
**创建时间**: {{FormatTime .Data.created_at}}
**帖子编号**: {{.Data.post_number}}
**话题 ID**: {{.Data.topic_id}}
**话题标题**: {{.Data.data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**原始帖子 ID**: {{.Data.data.original_post_id}}
**原始帖子类型**: {{.Data.data.original_post_type}}
**原始用户**: [{{.Header.Instance}}/u/{{.Data.data.original_username}}]({{.Header.Instance}}/u/{{.Data.data.original_username}})
**修订版本号**: {{.Data.data.revision_number}}
**显示用户**: [{{.Header.Instance}}/u/{{.Data.data.display_username}}]({{.Header.Instance}}/u/{{.Data.data.display_username}})
`,
	NotificationCreated6: `# Discourse 通知创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**通知标题**: {{.Data.fancy_title}}
**通知 ID**: {{.Data.id}}
**通知类型**: {{.Data.notification_type}}
**阅读状态**: {{.Data.read}}
**高优先级**: {{.Data.high_priority}}
**创建时间**: {{FormatTime .Data.created_at}}
**帖子编号**: {{.Data.post_number}}
**话题 ID**: {{.Data.topic_id}}
**话题标题**: {{.Data.data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**原始帖子 ID**: {{.Data.data.original_post_id}}
**原始帖子类型**: {{.Data.data.original_post_type}}
**原始用户**: [{{.Header.Instance}}/u/{{.Data.data.original_username}}]({{.Header.Instance}}/u/{{.Data.data.original_username}})
**修订版本号**: {{.Data.data.revision_number}}
**显示用户**: [{{.Header.Instance}}/u/{{.Data.data.display_username}}]({{.Header.Instance}}/u/{{.Data.data.display_username}})
**显示名称**: {{.Data.data.display_name}}
**群组名称**: {{.Data.data.group_name}}
`,
	NotificationCreated12: `# Discourse 通知创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**通知 ID**: {{.Data.id}}
**通知类型**: {{.Data.notification_type}}
**阅读状态**: {{.Data.read}}
**高优先级**: {{.Data.high_priority}}
**创建时间**: {{FormatTime .Data.created_at}}
**徽章 ID**: {{.Data.data.badge_id}}
**徽章名称**: {{.Data.data.badge_name}}
**徽章 Slug**: {{.Data.data.badge_slug}}
**徽章标题**: {{.Data.data.badge_title}}
**用户名**: [{{.Header.Instance}}/u/{{.Data.data.username}}]({{.Header.Instance}}/u/{{.Data.data.username}})
`,
	Ping: `# Discourse Ping 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**Ping**: {{.Data}}
`,
	PostCreated: `# Discourse 帖子创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子内容**: {{.Data.raw}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**创建时间**: {{FormatTime .Data.created_at}}
**作者 ID**: {{.Data.user_id}}
**作者姓名**: {{.Data.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
`,
	PostDestroyed: `# Discourse 帖子删除 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子内容**: {{.Data.raw}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**作者 ID**: {{.Data.user_id}}
**作者姓名**: {{.Data.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**删除时间**: {{FormatTime .Data.deleted_at}}
**删除人ID**: {{.Data.deleted_by.id}}
**删除人名称**: {{.Data.deleted_by.name}}
**删除人链接**: [{{.Header.Instance}}/u/{{.Data.deleted_by.username}}]({{.Header.Instance}}/u/{{.Data.deleted_by.username}})
`,
	PostEdited: `# Discourse 帖子编辑 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子内容**: {{.Data.raw}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**编辑时间**: {{FormatTime .Data.updated_at}}
**作者 ID**: {{.Data.user_id}}
**作者姓名**: {{.Data.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
`,
	ReviewableCreated: `# Discourse 审核创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**审核类型**: {{.Data.type}}
**类型来源**: {{.Data.type_source}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**目标类型**: {{.Data.target_type}}
**目标 ID**: {{.Data.target_id}}
**目标链接**: [{{.Data.target_url}}]({{.Data.target_url}})
**目标内容**: {{.Data.raw}}
**目标时间**: {{FormatTime .Data.target_created_at}}
**目标创建人ID**: {{.Data.target_created_by_id}}
**审核 ID**: {{.Data.id}}
**审核创建时间**: {{FormatTime .Data.created_at}}
`,
	ReviewableScoreUpdated: `# Discourse 审核分数更新 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**审核类型**: {{.Data.type}}
**类型来源**: {{.Data.type_source}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**目标类型**: {{.Data.target_type}}
**目标 ID**: {{.Data.target_id}}
**目标链接**: [{{.Data.target_url}}]({{.Data.target_url}})
**目标内容**: {{.Data.raw}}
**目标时间**: {{FormatTime .Data.target_created_at}}
**目标创建人ID**: {{.Data.target_created_by_id}}
**审核 ID**: {{.Data.id}}
**审核创建时间**: {{FormatTime .Data.created_at}}
**分数更新时间**: {{FormatTime .Data.post_updated_at}}
**当前分数**: {{.Data.score}}
`,
	ReviewableTransitionedTo: `# Discourse 审核状态转换 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**审核类型**: {{.Data.type}}
**类型来源**: {{.Data.type_source}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**目标类型**: {{.Data.target_type}}
**目标 ID**: {{.Data.target_id}}
**目标链接**: [{{.Data.target_url}}]({{.Data.target_url}})
**目标内容**: {{.Data.raw}}
**目标时间**: {{FormatTime .Data.target_created_at}}
**目标创建人ID**: {{.Data.target_created_by_id}}
**审核 ID**: {{.Data.id}}
**审核创建时间**: {{FormatTime .Data.created_at}}
**状态转换时间**: {{FormatTime .Data.post_updated_at}}
**当前状态**: {{.Data.status}}
`,
	TopicCreated: `# Discourse 话题创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.id}}]({{.Header.Instance}}/t/{{.Data.id}})
**创建时间**: {{FormatTime .Data.created_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者姓名**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
`,
	TopicEdited: `# Discourse 话题编辑 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.id}}]({{.Header.Instance}}/t/{{.Data.id}})
**创建时间**: {{FormatTime .Data.created_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者姓名**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
**帖子数量**: {{.Data.posts_count}}
**浏览次数**: {{.Data.views}}
`,
	UserAddedToGroup: `# Discourse 用户加入群组 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.user_id}}
**群组 ID**: {{.Data.group_id}}
**唯一 ID**: {{.Data.id}}
**通知级别**: {{.Data.notification_level}}
**加入时间**: {{FormatTime .Data.created_at}}
`,
	UserBadgeGranted: `# Discourse 用户徽章授予 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**徽章授予 ID**: {{.Data.id}}
**徽章 ID**: {{.Data.badge_id}}
**用户 ID**: {{.Data.user_id}}
**授予人 ID**: {{.Data.granted_by_id}}
**授予时间**: {{FormatTime .Data.granted_at}}
**创建时间**: {{FormatTime .Data.created_at}}
`,
	UserConfirmedEmail: `# Discourse 用户确认邮箱 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户姓名**: {{.Data.name}}
**用户链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**创建时间**: {{FormatTime .Data.created_at}}
`,
	UserCreated: `# Discourse 用户创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户姓名**: {{.Data.name}}
**用户链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**创建时间**: {{FormatTime .Data.created_at}}
`,
	UserLoggedIn: `# Discourse 用户登录 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户姓名**: {{.Data.name}}
**用户链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**登录时间**: {{FormatTime .Data.created_at}}
`,
	UserLoggedOut: `# Discourse 用户退出 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户姓名**: {{.Data.name}}
**用户链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**退出时间**: {{FormatTime .Data.created_at}}
`,
	UserUpdated: `# Discourse 用户更新 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户姓名**: {{.Data.name}}
**用户链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
`,
}
