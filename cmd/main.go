package main

import (
	"flag"
	"fmt"
	"log"
	"mdToHtml/internal/site"

	"github.com/spf13/viper"
)

func main() {
	build := flag.Bool("build", false, "清理当前根目录并从markdown生成")
	clean := flag.Bool("clean", false, "删除生成的html文件")
	help := flag.Bool("help", false, "查看帮助")
	init := flag.Bool("init", false, "根据配置文件初始化")
	flag.Parse()
	if *build {
		site.Build()
	} else if *clean {
		site.Clean()
	} else if *help {
		flag.Usage()
		fmt.Println("站点根目录位于wwwroot markdown根目录位于wwwmark")
	} else if *init {
		// 调用初始化 释放默认config.yaml
		site.Init()

		viper.SetConfigFile("config.yaml")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		webroot := viper.GetString("paths.web")
		markdownroot := viper.GetString("paths.markdown")
		template := viper.GetString("paths.template")
		fmt.Println(webroot)
		fmt.Println(markdownroot)
		fmt.Println(template)

		// 释放html模板 初始化markdown根目录
		site.InitHtmlTemplate(template)
		site.InitMdRoot(markdownroot)
	}
}
