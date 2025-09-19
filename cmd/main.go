package main

import (
	"flag"
	"fmt"
	"mdToHtml/internal/site"
)

func main() {
	build := flag.Bool("build", false, "清理当前根目录并从markdown生成")
	clean := flag.Bool("clean", false, "删除生成的html文件")
	help := flag.Bool("help", false, "查看帮助")
	flag.Parse()
	if *build {
		site.Build()
	} else if *clean {
		site.Clean()
	} else if *help {
		flag.Usage()
		fmt.Println("站点根目录位于wwwroot markdown根目录位于wwwmark")
	}
}
