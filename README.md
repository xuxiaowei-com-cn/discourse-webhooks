# Discourse WebHooks

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

| 事件类型  | 事件                   | 进度 | 说明     |
|-------|----------------------|----|--------|
| ping  | ping                 | ✅  | 测试联通性  |
| user  | user_created         | ✅  | 创建用户   |
| user  | user_confirmed_email | ✅  | 用户验证邮箱 |
| topic | topic_created        |    | 话题名字创建 |
| topic | topic_edited         |    | 话题名字修改 |
| post  | post_created         |    | 帖子内容创建 |
| post  | post_edited          |    | 帖子内容修改 |
