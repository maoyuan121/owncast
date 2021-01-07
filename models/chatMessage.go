package models

import (
	"bytes"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"mvdan.cc/xurls"
)

// ChatEvent 表示一条聊天消息
type ChatEvent struct {
	ClientID string `json:"-"` // client id

	Author      string    `json:"author,omitempty"`    // 作者
	Body        string    `json:"body,omitempty"`      // 内容
	ID          string    `json:"id"`                  // id
	MessageType string    `json:"type"`                // 消息类型
	Visible     bool      `json:"visible"`             // 是否可见
	Timestamp   time.Time `json:"timestamp,omitempty"` // 时间
}

// 检查消息是否合法
// 作者，内容，id 缺少其中一个便判断为非法
func (m ChatEvent) Valid() bool {
	return m.Author != "" && m.Body != "" && m.ID != ""
}

// 如果内容格式是 markdown，那么把 markdown 转换为 HTML，
// 并且对 HTML 进行消毒
func (m *ChatEvent) RenderAndSanitizeMessageBody() {
	raw := m.Body

	// Set the new, sanitized and rendered message body
	m.Body = RenderAndSanitize(raw)
}

// RenderAndSanitize 将 markdown 转为 HTML，并且消毒
func RenderAndSanitize(raw string) string {
	rendered := renderMarkdown(raw)
	safe := sanitize(rendered)

	// Set the new, sanitized and rendered message body
	return strings.TrimSpace(safe)
}

func renderMarkdown(raw string) string {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			extension.NewLinkify(
				extension.WithLinkifyAllowedProtocols([][]byte{
					[]byte("http:"),
					[]byte("https:"),
				}),
				extension.WithLinkifyURLRegexp(
					xurls.Strict,
				),
			),
		),
	)

	trimmed := strings.TrimSpace(raw)
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(trimmed), &buf); err != nil {
		panic(err)
	}

	return buf.String()
}

// 消毒
func sanitize(raw string) string {
	p := bluemonday.StrictPolicy()

	// Require URLs to be parseable by net/url.Parse
	p.AllowStandardURLs()

	// 允许 a 有 href 属性
	p.AllowAttrs("href").OnElements("a")

	// 将所有的 link 加上 noreferrer
	p.RequireNoReferrerOnLinks(true)

	// 将所有的 link 加上 target="_blank"
	p.AddTargetBlankToFullyQualifiedLinks(true)

	// 允许段落和换行
	p.AllowElements("br")
	p.AllowElements("p")

	// 允许图片有 src alt title 属性
	p.AllowElements("img")
	p.AllowAttrs("src").OnElements("img")
	p.AllowAttrs("alt").OnElements("img")
	p.AllowAttrs("title").OnElements("img")

	// Custom emoji have a class already specified.
	// We should only allow classes on emoji, not *all* imgs.
	// But TODO.
	p.AllowAttrs("class").OnElements("img")

	// 允许加粗
	p.AllowElements("strong")

	// 允许强调
	p.AllowElements("em")

	return p.Sanitize(raw)
}
