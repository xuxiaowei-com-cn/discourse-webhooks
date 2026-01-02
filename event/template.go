package event

var TemplateMap = map[Type]string{
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
**话题 ID**: {{.Data.topic_id}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**帖子 ID**: {{.Data.id}}
**帖子内容**: {{.Data.raw}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**创建时间**: {{FormatTime .Data.created_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者名称**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
`,
	PostDestroyed: `# Discourse 帖子删除 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题 ID**: {{.Data.topic_id}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**帖子 ID**: {{.Data.id}}
**帖子内容**: {{.Data.raw}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**作者 ID**: {{.Data.user_id}}
**作者名称**: {{.Data.name}}
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
**话题 ID**: {{.Data.topic_id}}
**话题标题**: {{.Data.topic_title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}]({{.Header.Instance}}/t/{{.Data.topic_id}})
**帖子 ID**: {{.Data.id}}
**帖子内容**: {{.Data.raw}}
**帖子链接**: [{{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}}]({{.Header.Instance}}/t/{{.Data.topic_id}}/{{.Data.post_number}})
**编辑时间**: {{FormatTime .Data.updated_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者名称**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
`,
	TopicCreated: `# Discourse 话题创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题 ID**: {{.Data.id}}
**话题标题**: {{.Data.title}}
**话题链接**: [{{.Header.Instance}}/t/{{.Data.id}}]({{.Header.Instance}}/t/{{.Data.id}})
**创建时间**: {{FormatTime .Data.created_at}}
**作者 ID**: {{.Data.created_by.id}}
**作者名称**: {{.Data.created_by.name}}
**作者链接**: [{{.Header.Instance}}/u/{{.Data.created_by.username}}]({{.Header.Instance}}/u/{{.Data.created_by.username}})
`,
	TopicEdited: `# Discourse 话题编辑 事件通知
**实例地址**: {{.Header.Instance}}
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**话题信息**:
- 话题 ID: {{.Data.ID}}
- 话题标题: {{.Data.Title}}
- 创建时间: {{FormatTime .Data.CreatedAt}}
- 创建者: {{.Data.CreatedByUsername}}
`,
	UserCreated: `# Discourse 用户创建 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户名称**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**创建时间**: {{FormatTime .Data.created_at}}
`,
	UserAddedToGroup: `# Discourse 用户加入群组 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**群组 ID**: {{.Data.group_id}}
**唯一 ID**: {{.Data.id}}
**通知级别**: {{.Data.notification_level}}
**创建时间**: {{FormatTime .Data.created_at}}
`,
	UserConfirmedEmail: `# Discourse 用户确认邮箱 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户名称**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**创建时间**: {{FormatTime .Data.created_at}}
`,
	UserLoggedIn: `# Discourse 用户登录 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户名称**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**登录时间**: {{FormatTime .Data.created_at}}
`,
	UserLoggedOut: `# Discourse 用户退出 事件通知
**实例地址**: [{{.Header.Instance}}]({{.Header.Instance}})
**事件 ID**: {{.Header.EventId}}
**事件类型**: {{.Header.Event}}
**用户 ID**: {{.Data.id}}
**用户名称**: [{{.Header.Instance}}/u/{{.Data.username}}]({{.Header.Instance}}/u/{{.Data.username}})
**用户邮箱**: {{.Data.email}}
**退出时间**: {{FormatTime .Data.created_at}}
`,
}
