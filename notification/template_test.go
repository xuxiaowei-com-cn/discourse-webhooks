package notification

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xuxiaowei-com-cn/discourse-webhooks/event"
)

// TestRenderTemplateWithSampleData 使用 data 目录下的真实 Webhook 请求体渲染模板，
// 确保模板不会因为请求体中的 null 字段（如 deleted_by）而执行失败。
func TestRenderTemplateWithSampleData(t *testing.T) {
	bodyFiles, err := filepath.Glob("../data/*/*/*body*.json")
	if err != nil {
		t.Fatalf("glob data files: %v", err)
	}
	if len(bodyFiles) == 0 {
		t.Fatal("no sample webhook body files found")
	}

	for _, bodyFile := range bodyFiles {
		t.Run(bodyFile, func(t *testing.T) {
			dir := filepath.Dir(bodyFile)
			headerFiles, err := filepath.Glob(filepath.Join(dir, "*header*.json"))
			if err != nil || len(headerFiles) == 0 {
				t.Skipf("no header file in %s", dir)
			}

			headerBytes, err := os.ReadFile(headerFiles[0])
			if err != nil {
				t.Fatalf("read header: %v", err)
			}
			var header event.Discourse
			if err := json.Unmarshal(headerBytes, &header); err != nil {
				t.Fatalf("unmarshal header: %v", err)
			}

			bodyBytes, err := os.ReadFile(bodyFile)
			if err != nil {
				t.Fatalf("read body: %v", err)
			}
			var payload map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &payload); err != nil {
				t.Fatalf("unmarshal body: %v", err)
			}

			templateContent := templateFor(header, payload[header.EventType])
			if templateContent == "" {
				t.Skipf("no template for event %s", header.Event)
			}

			out, err := renderTemplate(templateContent, header, payload[header.EventType])
			if err != nil {
				t.Fatalf("render template for %s: %v", header.Event, err)
			}
			if strings.TrimSpace(out) == "" {
				t.Fatalf("rendered template for %s is empty", header.Event)
			}
		})
	}
}

// TestRenderDestroyedEventWithNullDeletedBy 复现 deleted_by 为 null 时模板渲染失败的问题：
// Discourse 删除话题/帖子时可能不记录删除人（deleted_by 为 null），
// 模板不能因此返回 “nil pointer evaluating interface {}.id” 而中断整条通知。
func TestRenderDestroyedEventWithNullDeletedBy(t *testing.T) {
	header := event.Discourse{
		Instance:  "https://xuxiaowei.io",
		EventId:   "30613",
		EventType: "topic",
	}
	cases := []struct {
		name  string
		tmpl  event.Type
		event string
		data  map[string]interface{}
	}{
		{
			name:  "topic_destroyed",
			tmpl:  event.TopicDestroyed,
			event: "topic_destroyed",
			data: map[string]interface{}{
				"id":                1498,
				"title":             "华为应用市场\tHarmonyOS",
				"created_at":        "2026-11-02T11:48:54.871Z",
				"deleted_at":        nil,
				"deleted_by":        nil,
				"created_by":        map[string]interface{}{"id": 268, "username": "fuyang", "name": "ffffffff"},
				"last_poster":       map[string]interface{}{"id": 268, "username": "fuyang", "name": "ffffffff"},
				"last_posted_at":    "2026-11-02T11:48:54.929Z",
				"posts_count":       1,
				"views":             0,
				"reply_count":       0,
				"like_count":        0,
				"word_count":        5,
				"participant_count": 1,
				"tags":              []interface{}{},
				"tags_descriptions": map[string]interface{}{},
			},
		},
		{
			name:  "post_destroyed",
			tmpl:  event.PostDestroyed,
			event: "post_destroyed",
			data: map[string]interface{}{
				"id":          10,
				"topic_id":    9,
				"topic_title": "测试文件上传阿里云对象储存 oss",
				"user_id":     1,
				"name":        "徐晓伟",
				"username":    "xuxiaowei",
				"post_number": 1,
				"posts_count": 0,
				"raw":         "![猫屎二号|500x500](upload://ohGFmibN4slI3lTHheqRzJy3ils.jpeg)",
				"deleted_at":  nil,
				"deleted_by":  nil,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			header.Event = c.event
			out, err := renderTemplate(event.TemplateMap[c.tmpl], header, c.data)
			if err != nil {
				t.Fatalf("render %s with null deleted_by: %v", c.event, err)
			}
			if strings.Contains(out, "删除人 ID") || strings.Contains(out, "删除人ID") {
				t.Fatalf("deleted_by is null, but deleter fields were rendered:\n%s", out)
			}
			if !strings.Contains(out, "**删除人**: 未知") {
				t.Fatalf("missing fallback text for null deleted_by:\n%s", out)
			}
			if strings.Contains(out, "\n\n") {
				t.Fatalf("rendered %s contains unexpected blank lines:\n%s", c.event, out)
			}
		})
	}
}
