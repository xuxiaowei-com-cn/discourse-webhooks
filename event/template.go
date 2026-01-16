package event

var TemplateMap = map[Type]string{
	AcceptedSolution: `# Discourse 接受解决方案 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**话题 ID**: {{.Data.topic_id}}
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子编号**: {{.Data.post_number}}
**帖子链接**: [{{.Header.Instance}}{{.Data.post_url}}]({{.Header.Instance}}{{.Data.post_url}})
**帖子内容**: {{LimitLength .Data.raw}}
**用户 ID**: {{.Data.user_id}}
**用户姓名**: {{.Data.name}}
**用户链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**创建时间**: {{FormatTime .Data.created_at}}
**更新时间**: {{FormatTime .Data.updated_at}}
`,
	CategoryCreated: `# Discourse 分类创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**分类 ID**: {{.Data.id}}
**分类名称**: {{.Data.name}}
**分类 Slug**: {{.Data.slug}}
**分类链接**: [{{.Header.Instance}}{{.Data.topic_url}}]({{.Header.Instance}}{{.Data.topic_url}})
**话题数量**: {{.Data.topic_count}}
**帖子数量**: {{.Data.post_count}}
**分类颜色**: {{.Data.color}}
**文字颜色**: {{.Data.text_color}}
**分类位置**: {{.Data.position}}
**父分类 ID**: {{.Data.parent_category_id}}
**分类描述**: {{.Data.description}}
**是否限制阅读**: {{.Data.read_restricted}}
**是否允许徽章**: {{.Data.allow_badges}}
**是否允许话题特色链接**: {{.Data.topic_featured_link_allowed}}
`,
	CategoryUpdated: `# Discourse 分类更新 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**分类 ID**: {{.Data.id}}
**分类名称**: {{.Data.name}}
**分类 Slug**: {{.Data.slug}}
**分类链接**: [{{.Header.Instance}}{{.Data.topic_url}}]({{.Header.Instance}}{{.Data.topic_url}})
**话题数量**: {{.Data.topic_count}}
**帖子数量**: {{.Data.post_count}}
**分类颜色**: {{.Data.color}}
**文字颜色**: {{.Data.text_color}}
**分类位置**: {{.Data.position}}
**分类描述**: {{.Data.description}}
**是否限制阅读**: {{.Data.read_restricted}}
**是否允许徽章**: {{.Data.allow_badges}}
**是否允许话题特色链接**: {{.Data.topic_featured_link_allowed}}
**无法删除原因**: {{.Data.cannot_delete_reason}}
`,
	GroupUpdated: `# Discourse 群组更新 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**群组 ID**: {{.Data.id}}
**群组名称**: {{.Data.name}}
**显示名称**: {{.Data.display_name}}
**用户数量**: {{.Data.user_count}}
**是否自动**: {{.Data.automatic}}
**可见性级别**: {{.Data.visibility_level}}
**是否主要群组**: {{.Data.primary_group}}
**简介摘要**: {{.Data.bio_excerpt}}
**是否允许成员请求**: {{.Data.allow_membership_requests}}
**默认通知级别**: {{.Data.default_notification_level}}
**是否可以查看成员**: {{.Data.can_see_members}}
**是否可以管理群组**: {{.Data.can_admin_group}}
`,
	NotificationCreated1: `# Discourse 通知创建 事件通知
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
**话题链接**: [{{.Header.Instance}}/t/{{.Data.slug}}/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.slug}}/{{.Data.topic_id}})
**原始帖子 ID**: {{.Data.data.original_post_id}}
**原始帖子类型**: {{.Data.data.original_post_type}}
**原始用户**: [{{.Header.Instance}}/u/{{.Data.data.original_username}}]({{.Header.Instance}}/u/{{.Data.data.original_username}})
**修订版本号**: {{.Data.data.revision_number}}
**显示用户**: [{{.Header.Instance}}/u/{{.Data.data.display_username}}]({{.Header.Instance}}/u/{{.Data.data.display_username}})
`,
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
**显示信息**: {{.Data.data.display_username}}
`,
	NotificationCreated4: `# Discourse 通知创建 事件通知
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
	NotificationCreated5: `# Discourse 通知创建 事件通知
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
	NotificationCreated14: `# Discourse 通知创建 事件通知
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
**消息**: {{.Data.data.message}}
**标题**: {{.Data.data.title}}
**显示用户**: [{{.Header.Instance}}/u/{{.Data.data.display_username}}]({{.Header.Instance}}/u/{{.Data.data.display_username}})
`,
	NotificationCreated25: `# Discourse 通知创建 事件通知
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
**话题链接**: [{{.Header.Instance}}/t/{{.Data.slug}}/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.slug}}/{{.Data.topic_id}})
**原始帖子 ID**: {{.Data.data.original_post_id}}
**原始帖子类型**: {{.Data.data.original_post_type}}
**原始用户**: [{{.Header.Instance}}/u/{{.Data.data.original_username}}]({{.Header.Instance}}/u/{{.Data.data.original_username}})
**修订版本号**: {{.Data.data.revision_number}}
**显示用户**: [{{.Header.Instance}}/u/{{.Data.data.display_username}}]({{.Header.Instance}}/u/{{.Data.data.display_username}})
**显示名称**: {{.Data.data.display_name}}
**反应图标**: {{.Data.data.reaction_icon}}
`,
	NotificationCreated38: `# Discourse 通知创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**通知 ID**: {{.Data.id}}
**通知类型**: {{.Data.notification_type}}
**阅读状态**: {{.Data.read}}
**高优先级**: {{.Data.high_priority}}
**创建时间**: {{FormatTime .Data.created_at}}
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
**创建时间**: {{FormatTime .Data.created_at}}
**作者 ID**: {{.Data.user_id}}
**作者姓名**: {{.Data.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**帖子内容**: {{LimitLength .Data.raw}}
`,
	PostDestroyed: `# Discourse 帖子删除 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**作者 ID**: {{.Data.user_id}}
**作者姓名**: {{.Data.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**删除时间**: {{FormatTime .Data.deleted_at}}
**删除人ID**: {{.Data.deleted_by.id}}
**删除人名称**: {{.Data.deleted_by.name}}
**删除人链接**: [{{.Header.Instance}}/u/{{.Data.deleted_by.username}}]({{.Header.Instance}}/u/{{.Data.deleted_by.username}})
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**帖子内容**: {{LimitLength .Data.raw}}
`,
	PostEdited: `# Discourse 帖子编辑 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**编辑时间**: {{FormatTime .Data.updated_at}}
**作者 ID**: {{.Data.user_id}}
**作者姓名**: {{.Data.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**帖子内容**: {{LimitLength .Data.raw}}
`,
	PostLiked: `# Discourse 帖子点赞 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题 ID**: {{.Data.post.topic_id}}
**话题标题**: {{.Data.post.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.post.topic_id}}]({{.Header.Instance}}/t/{{.Data.post.topic_id}})
**帖子数量**: {{.Data.post.posts_count}}
**点赞用户 ID**: {{.Data.user.id}}
**点赞用户名**: [{{.Header.Instance}}/u/{{.Data.user.username}}]({{.Header.Instance}}/u/{{.Data.user.username}})
**点赞用户姓名**: {{.Data.user.name}}
**创建时间**: {{FormatTime .Data.post.created_at}}
**更新时间**: {{FormatTime .Data.post.updated_at}}
**回复数量**: {{.Data.post.reply_count}}
**阅读次数**: {{.Data.post.reads}}
**评   分**: {{.Data.post.score}}
**帖子 ID**: {{.Data.post.id}}
**帖子编号**: {{.Data.post.post_number}}
**帖子类型**: {{.Data.post.post_type}}
**帖子作者 ID**: {{.Data.post.user_id}}
**帖子作者姓名**: {{.Data.post.name}}
**帖子作者链接**: [{{.Header.Instance}}/u/{{.Data.post.username}}]({{.Header.Instance}}/u/{{.Data.post.username}})
**帖子链接**: [{{.Header.Instance}}{{.Data.post.post_url}}]({{.Header.Instance}}{{.Data.post.post_url}})
**帖子内容**: {{LimitLength .Data.post.raw}}
`,
	PostRecovered: `# Discourse 帖子恢复 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**创建时间**: {{FormatTime .Data.created_at}}
**更新时间**: {{FormatTime .Data.updated_at}}
**作者 ID**: {{.Data.user_id}}
**作者姓名**: {{.Data.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**帖子数量**: {{.Data.posts_count}}
**帖子 ID**: {{.Data.id}}
**帖子链接**: [{{.Header.Instance}}{{.Data.post_url}}]({{.Header.Instance}}{{.Data.post_url}})
**帖子内容**: {{LimitLength .Data.raw}}
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
**目标内容**: {{LimitLength .Data.raw}}
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
**目标内容**: {{LimitLength .Data.raw}}
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
**目标内容**: {{LimitLength .Data.raw}}
**目标时间**: {{FormatTime .Data.target_created_at}}
**目标创建人ID**: {{.Data.target_created_by_id}}
**审核 ID**: {{.Data.id}}
**审核创建时间**: {{FormatTime .Data.created_at}}
**状态转换时间**: {{FormatTime .Data.post_updated_at}}
**当前状态**: {{.Data.status}}
`,
	TagCreated: `# Discourse 标签创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**标签 ID**: {{.Data.id}}
**标签名称**: {{.Data.name}}
**话题数量**: {{.Data.topic_count}}
**人员标签**: {{.Data.staff}}
**标签描述**: {{.Data.description}}
`,
	TagDestroyed: `# Discourse 标签删除 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**标签 ID**: {{.Data.id}}
**标签名称**: {{.Data.name}}
**话题数量**: {{.Data.topic_count}}
**人员标签**: {{.Data.staff}}
**标签描述**: {{.Data.description}}
`,
	TopicClosedStatusUpdated: `# Discourse 话题关闭状态更新 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.id}}]({{.Header.Instance}}/t/{{.Data.id}})
**话题 ID**: {{.Data.id}}
**关闭状态**: {{.Data.closed}}
**创建时间**: {{FormatTime .Data.created_at}}
**最后回复时间**: {{FormatTime .Data.last_posted_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者姓名**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
**最后回复者 ID**: {{.Data.last_poster.id}}
**最后回复者姓名**: {{.Data.last_poster.name}}
**最后回复者链接**: [{{.Header.Instance}}/u/{{.Data.last_poster.username}}]({{.Header.Instance}}/u/{{.Data.last_poster.username}})
**帖子数量**: {{.Data.posts_count}}
**浏览次数**: {{.Data.views}}
**回复数量**: {{.Data.reply_count}}
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
	TopicDestroyed: `# Discourse 话题删除 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.id}}]({{.Header.Instance}}/t/{{.Data.id}})
**话题 ID**: {{.Data.id}}
**创建时间**: {{FormatTime .Data.created_at}}
**删除时间**: {{FormatTime .Data.deleted_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者姓名**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
**删除人 ID**: {{.Data.deleted_by.id}}
**删除人姓名**: {{.Data.deleted_by.name}}
**删除人链接**: [{{.Header.Instance}}/u/{{.Data.deleted_by.username}}]({{.Header.Instance}}/u/{{.Data.deleted_by.username}})
**最后回复者 ID**: {{.Data.last_poster.id}}
**最后回复者姓名**: {{.Data.last_poster.name}}
**最后回复者链接**: [{{.Header.Instance}}/u/{{.Data.last_poster.username}}]({{.Header.Instance}}/u/{{.Data.last_poster.username}})
**最后回复时间**: {{FormatTime .Data.last_posted_at}}
**帖子数量**: {{.Data.posts_count}}
**浏览次数**: {{.Data.views}}
**回复数量**: {{.Data.reply_count}}
**点赞数量**: {{.Data.like_count}}
**字数**: {{.Data.word_count}}
**参与人数**: {{.Data.participant_count}}
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
	TopicPinnedStatusUpdated: `# Discourse 话题置顶状态更新 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.id}}]({{.Header.Instance}}/t/{{.Data.id}})
**话题 ID**: {{.Data.id}}
**置顶状态**: {{.Data.pinned}}
**置顶类型**: {{.Data.pinned_globally}}
**创建时间**: {{FormatTime .Data.created_at}}
**最后回复时间**: {{FormatTime .Data.last_posted_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者姓名**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
**最后回复者 ID**: {{.Data.last_poster.id}}
**最后回复者姓名**: {{.Data.last_poster.name}}
**最后回复者链接**: [{{.Header.Instance}}/u/{{.Data.last_poster.username}}]({{.Header.Instance}}/u/{{.Data.last_poster.username}})
**帖子数量**: {{.Data.posts_count}}
**浏览次数**: {{.Data.views}}
**回复数量**: {{.Data.reply_count}}
`,
	TopicRecovered: `# Discourse 话题恢复 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题标题**: {{.Data.title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.id}}]({{.Header.Instance}}/t/{{.Data.id}})
**话题 ID**: {{.Data.id}}
**创建时间**: {{FormatTime .Data.created_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者姓名**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
**最后回复者 ID**: {{.Data.last_poster.id}}
**最后回复者姓名**: {{.Data.last_poster.name}}
**最后回复者链接**: [{{.Header.Instance}}/u/{{.Data.last_poster.username}}]({{.Header.Instance}}/u/{{.Data.last_poster.username}})
**最后回复时间**: {{FormatTime .Data.last_posted_at}}
**帖子数量**: {{.Data.posts_count}}
**浏览次数**: {{.Data.views}}
**回复数量**: {{.Data.reply_count}}
**点赞数量**: {{.Data.like_count}}
**字数**: {{.Data.word_count}}
**参与人数**: {{.Data.participant_count}}
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
	UserBadgeRevoked: `# Discourse 用户徽章撤销 事件通知
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
	UserDestroyed: `# Discourse 用户删除 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户姓名**: {{.Data.name}}
**用户链接**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**创建时间**: {{FormatTime .Data.created_at}}
**最后活跃时间**: {{FormatTime .Data.last_seen_at}}
**帖子数量**: {{.Data.post_count}}
**徽章数量**: {{.Data.badge_count}}
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
