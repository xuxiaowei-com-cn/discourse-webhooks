PROJECT := discourse-webhooks
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GIT_TAG := $(shell git describe --tags --always --dirty 2>/dev/null || echo $(GIT_COMMIT))
BUILD_TIME := $(shell TZ='Asia/Shanghai' date '+%Y-%m-%dT%H:%M:%S%z')
CiPipelineUrl := "$$CI_PIPELINE_URL"
# 尽量减小体积：剔除符号表与调试信息、移除编译路径、禁用内置 VCS 元数据、使用纯 Go 网络库
GOFLAGS := -trimpath -buildvcs=false
TAGS := netgo
LDFLAGS := -s -w -X main.GitCommit=$(GIT_COMMIT) -X main.GitBranch=$(GIT_BRANCH) -X main.GitTag=$(GIT_TAG) -X main.BuildTime=$(BUILD_TIME) -X main.CiPipelineUrl=$(CiPipelineUrl)

.PHONY: all build run version clean release linux darwin windows

all: build

build:
	CGO_ENABLED=0 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)

run:
	CGO_ENABLED=0 go run $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" main.go

version:
	CGO_ENABLED=0 go run $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" main.go --version

clean:
	rm -f $(PROJECT) $(PROJECT)-* main

release: linux darwin windows

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)-linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)-linux-arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=loong64 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)-linux-loong64

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)-darwin-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)-darwin-arm64

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build $(GOFLAGS) -tags "$(TAGS)" -ldflags "$(LDFLAGS)" -o $(PROJECT)-windows-arm64.exe

.PHONY: compress
compress:
	@command -v upx >/dev/null 2>&1 && upx --lzma -9 $(PROJECT) $(PROJECT)-* || echo "upx not found, skip compress"
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(PROJECT)-windows-arm64.exe
