package site

import (
	"fmt"
	"log"
	"os"
)

const indexMd = `## 主页`

const testMd = `## 测试页面`

const styleCss = `
/* --- 全局样式 --- */
body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
    line-height: 1.6;
    color: #24292e;
    background-color: #ffffff;
    padding: 20px;
}

/* --- 内容布局 --- */
body {
    max-width: 880px;
    margin: 40px auto;
    padding: 0 30px;
}

/* --- 标题 --- */
h1, h2, h3, h4, h5, h6 {
    margin-top: 24px;
    margin-bottom: 16px;
    font-weight: 600;
    line-height: 1.25;
    border-bottom: 1px solid #eaecef;
    padding-bottom: 0.3em;
}
h1 { font-size: 2em; }
h2 { font-size: 1.5em; }
h3 { font-size: 1.25em; }

/* --- 基础元素 --- */
p { margin-bottom: 16px; }
hr {
    height: 0.25em;
    padding: 0;
    margin: 24px 0;
    background-color: #e1e4e8;
    border: 0;
}
a {
    color: #0366d6;
    text-decoration: none;
}
a:hover { text-decoration: underline; }
img { max-width: 100%; height: auto; }

/* --- 列表 --- */
ul, ol {
    padding-left: 2em;
    margin-bottom: 16px;
}
/* 任务列表 */
li.task-list-item {
    list-style-type: none;
}
input[type="checkbox"] {
    margin-right: 8px;
}

/* --- 代码块 --- */
pre {
    background-color: #f6f8fa;
    border-radius: 6px;
    padding: 16px;
    overflow: auto;
    font-size: 85%;
    line-height: 1.45;
}
code {
    font-family: "SFMono-Regular", Consolas, "Liberation Mono", Menlo, Courier, monospace;
}
/* 行内代码 */
p > code, li > code {
    padding: 0.2em 0.4em;
    margin: 0;
    font-size: 85%;
    background-color: rgba(27,31,35,0.05);
    border-radius: 3px;
}
pre > code {
    padding: 0;
    margin: 0;
    background-color: transparent;
    border: 0;
}

/* --- 引用 --- */
blockquote {
    padding: 0 1em;
    color: #6a737d;
    border-left: 0.25em solid #dfe2e5;
    margin-left: 0;
    margin-right: 0;
    margin-bottom: 16px;
}

/* --- 表格 --- */
table {
    width: 100%;
    border-collapse: collapse;
    margin-bottom: 16px;
    display: block;
    overflow: auto;
}
th, td {
    padding: 6px 13px;
    border: 1px solid #dfe2e5;
}
th {
    font-weight: 600;
    background-color: #f6f8fa;
}
tr {
    background-color: #fff;
    border-top: 1px solid #c6cbd1;
}
tr:nth-child(2n) {
    background-color: #f6f8fa;
}

/* --- 脚注 --- */
.footnotes {
    margin-top: 3em;
    padding-top: 1em;
    border-top: 1px solid #eaecef;
    font-size: 0.9em;
    color: #6a737d;
}`

const pagehtmlHeader = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nsite/page</title>
    <link rel="stylesheet" href="../style.css">
</head>
<body>
<div id="main-header" style="background-color: white; box-shadow: 0 2px 5px rgba(0,0,0,0.1); width: 100%; position: fixed; top: 0; left: 0; z-index: 1000;">
    <div style="max-width: 1100px; margin: 0 auto; padding: 0 20px; height: 65px; display: flex; justify-content: space-between; align-items: center;">
        
        <a href="/" style="font-size: 1.6em; font-weight: bold; color: #333; text-decoration: none;">
            Nsite
        </a>

        <div style="display: flex; gap: 25px;">
            <a href="/index.html" style="color: #555; text-decoration: none; font-size: 1.1em;">首页</a>
            <a href="/categories.html" style="color: #555; text-decoration: none; font-size: 1.1em;">分类</a>
            <a href="/tags.html" style="color: #555; text-decoration: none; font-size: 1.1em;">标签</a>
            <a href="/about.html" style="color: #555; text-decoration: none; font-size: 1.1em;">关于我</a>
        </div>
    </div>
</div>
<div class="header-spacer" style="height: 30px;"></div>
`

const indexhtmlHeader = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nsite/index</title>
    <link rel="stylesheet" href="./style.css">
</head>
<body>
<div id="main-header" style="background-color: white; box-shadow: 0 2px 5px rgba(0,0,0,0.1); width: 100%; position: fixed; top: 0; left: 0; z-index: 1000;">
    <div style="max-width: 1100px; margin: 0 auto; padding: 0 20px; height: 65px; display: flex; justify-content: space-between; align-items: center;">
        
        <a href="/" style="font-size: 1.6em; font-weight: bold; color: #333; text-decoration: none;">
            Nsite
        </a>

        <div style="display: flex; gap: 25px;">
            <a href="/index.html" style="color: #555; text-decoration: none; font-size: 1.1em;">首页</a>
            <a href="/categories.html" style="color: #555; text-decoration: none; font-size: 1.1em;">分类</a>
            <a href="/tags.html" style="color: #555; text-decoration: none; font-size: 1.1em;">标签</a>
            <a href="/about.html" style="color: #555; text-decoration: none; font-size: 1.1em;">关于我</a>
        </div>
    </div>
</div>
<div class="header-spacer" style="height: 30px;"></div>`

const htmlEnd = `
</br>
</br>
<!-- mathjax 公式 -->
<script src="https://cdn.jsdelivr.net/npm/mathjax@4/tex-mml-chtml.js" defer></script> 
<!-- prismjs 代码高亮 -->
<script src="https://cdn.jsdelivr.net/npm/prismjs@1.30.0/prism.min.js"></script>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/prismjs@1.30.0/themes/prism.min.css" integrity="sha256-ko4j5rn874LF8dHwW29/xabhh8YBleWfvxb8nQce4Fc=" crossorigin="anonymous">

<footer style="bottom: 0; left: 0; width: 100%; text-align: center; padding: 10px 0; font-size: 0.9em; color: #555;">
    Powered by <a href="https://example.com" target="_blank" style="color: #555; text-decoration: none;">Nsite</a>
</footer>
</body>`

func Init() {
	configTemplate := `
paths:
  web: "wwwroot"
  markdown: "wwwmark"
  template: "template"
`
	configFile, err := os.Create("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	writeString, err := configFile.WriteString(configTemplate)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(writeString)
}

func InitHtmlTemplate(htmlTemplateDir string) {
	err := os.MkdirAll(htmlTemplateDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// 创建文件
	indexHtmlHeaderFile, err := os.Create(htmlTemplateDir + "/indexHtmlHeader.html")
	if err != nil {
		log.Fatal(err)
	}
	defer indexHtmlHeaderFile.Close()

	pageHtmlHeaderFile, err := os.Create(htmlTemplateDir + "/pageHtmlHeader.html")
	if err != nil {
		log.Fatal(err)
	}
	defer pageHtmlHeaderFile.Close()

	htmlEndFile, err := os.Create(htmlTemplateDir + "/htmlEnd.html")
	if err != nil {
		log.Fatal(err)
	}
	defer htmlEndFile.Close()

	//写入文件
	_, _ = indexHtmlHeaderFile.WriteString(indexhtmlHeader)
	_, _ = pageHtmlHeaderFile.WriteString(pagehtmlHeader)
	_, _ = htmlEndFile.WriteString(htmlEnd)

}

func InitMdRoot(mdRootDir string) {
	err := os.MkdirAll(mdRootDir+"/page", 0755)
	if err != nil {
		log.Fatal(err)
	}

	indexMdFile, err := os.Create(mdRootDir + "/index.md")
	if err != nil {
		log.Fatal(err)
	}
	defer indexMdFile.Close()

	testMdFile, err := os.Create(mdRootDir + "/page/test.md")
	if err != nil {
		log.Fatal(err)
	}
	defer testMdFile.Close()

	styleCssFile, err := os.Create(mdRootDir + "/style.css")
	if err != nil {
		log.Fatal(err)
	}
	defer styleCssFile.Close()

	_, _ = indexMdFile.WriteString(indexMd)
	_, _ = testMdFile.WriteString(testMd)
	_, _ = styleCssFile.WriteString(styleCss)

}
