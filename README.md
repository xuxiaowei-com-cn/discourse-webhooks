# Discourse WebHooks：Discourse 论坛 WebHooks

## 开发命令

### get

```shell
go env -w GOPROXY=https://goproxy.cn,direct
# go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct
# go env -w GOPROXY=https://goproxy.io,direct
go get -u github.com/urfave/cli/v3
```

```shell
go get -u all
```

### mod

```shell
go mod download
```

```shell
go mod tidy
```

### test

```shell
go test ./... -v
```

## build

```shell
# 简化
go build main.go
```

```shell
# 完整
make release
```

## 实现

| 事件类型         | 事件                          | 进度 | 说明       |
|--------------|-----------------------------|----|----------|
| ping         | ping                        | ✅  | 测试联通性    |
| solved       | accepted_solution           | ✅  | 接受解决方案   |
| category     | category_created            | ✅  | 分类创建     |
| category     | category_updated            | ✅  | 分类更新     |
| group        | group_updated               | ✅  | 群组更新     |
| notification | notification_created1       | ✅  | 通知创建     |
| notification | notification_created2       | ✅  | 通知创建     |
| notification | notification_created4       | ✅  | 通知创建     |
| notification | notification_created5       | ✅  | 通知创建     |
| notification | notification_created6       | ✅  | 通知创建     |
| notification | notification_created12      | ✅  | 通知创建     |
| notification | notification_created14      | ✅  | 通知创建     |
| notification | notification_created25      | ✅  | 通知创建     |
| notification | notification_created38      | ✅  | 通知创建     |
| post         | post_created                | ✅  | 帖子内容创建   |
| post         | post_destroyed              | ✅  | 帖子删除     |
| post         | post_edited                 | ✅  | 帖子编辑     |
| like         | post_liked                  | ✅  | 帖子点赞     |
| post         | post_recovered              | ✅  | 帖子恢复     |
| reviewable   | reviewable_created          | ✅  | 审核创建     |
| reviewable   | reviewable_score_updated    | ✅  | 审核分数更新   |
| reviewable   | reviewable_transitioned_to  | ✅  | 审核状态转换   |
| tag          | tag_created                 | ✅  | 标签创建     |
| topic        | topic_closed_status_updated | ✅  | 话题关闭状态更新 |
| topic        | topic_created               | ✅  | 话题创建     |
| topic        | topic_destroyed             | ✅  | 话题删除     |
| topic        | topic_edited                | ✅  | 话题编辑     |
| topic        | topic_pinned_status_updated | ✅  | 话题置顶状态更新 |
| topic        | topic_recovered             | ✅  | 话题恢复     |
| group_user   | user_added_to_group         | ✅  | 用户加入群组   |
| user_badge   | user_badge_granted          | ✅  | 用户徽章授予   |
| user_badge   | user_badge_revoked          | ✅  | 用户徽章撤销   |
| user         | user_confirmed_email        | ✅  | 用户确认邮箱   |
| user         | user_created                | ✅  | 用户创建     |
| user         | user_logged_in              | ✅  | 用户登录     |
| user         | user_logged_out             | ✅  | 用户退出     |
| user         | user_updated                | ✅  | 用户更新     |

## 联系我

![wechat-work.jpg](static/wechat-work.jpg)
