package site

import (
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func MdFull() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.TaskList,
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
			extension.Typographer,
			extension.CJK,

			mathjax.MathJax, // 数学公式支持
		),
		// 唯一id
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		// 配置 HTML 渲染器
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)
}
