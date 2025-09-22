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

	// 设置config文件路径
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if *build {
		logo()
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("config不存在或系统未初始化，请使用-init初始化")
			log.Fatal(err)

		}
		webroot := viper.GetString("paths.web")
		markdownroot := viper.GetString("paths.markdown")
		template := viper.GetString("paths.template")

		site.Build(webroot, markdownroot, template)
	} else if *clean {
		webroot := viper.GetString("paths.web")
		site.Clean(webroot)
	} else if *help {
		logo()
		flag.Usage()
		fmt.Println("站点根目录位于wwwroot markdown根目录位于wwwmark")
	} else if *init {
		logo()
		site.Init()
		// 调用初始化 释放默认config.yaml 与相关模板
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

func logo() {
	logo := `
   _  __  ____  ____ ______  ____
  / |/ / / __/ /  _//_  __/ / __/
 /    / _\ \  _/ /   / /   / _/  
/_/|_/ /___/ /___/  /_/   /___/  
`
	fmt.Println(logo)
}
