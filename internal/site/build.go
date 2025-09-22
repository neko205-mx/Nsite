package site

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// markdown 定义html文件和markdown文件位置的结构体
type markdown struct {
	targetFile  string
	pointToFile string
}

// Build 遍历 markdownPath 目录下的文件，处理所有 .css 文件和包含 "page" 的 .md 文件，
// 将它们复制或转换为对应的 .html 文件到 webPath 目录下。
// 同时查找并处理 index.md 文件作为网站的首页。
func Build(webPath, markdownPath, templatePath string) {
	fmt.Println("building...")
	Clean(webPath)
	err := filepath.WalkDir(markdownPath, func(path string, d os.DirEntry, err error) error {
		fmt.Println("Processing file: " + path)
		if strings.Contains(path, ".css") {
			fileCopy(path, markdownPath, webPath)
		}
		if err != nil {
			return err
		}
		if strings.Contains(path, "page") && strings.Contains(path, "md") {
			targetFile := path
			pointToFile := strings.Replace(targetFile, markdownPath, webPath, 1)
			pointToFile = strings.Replace(pointToFile, "md", "html", 1)

			mdFile := markdown{
				targetFile:  targetFile,
				pointToFile: pointToFile,
			}
			fileProcessing(mdFile, templatePath)
		}
		//fmt.Printf("%s\n", path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(markdownPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "index.md") {
			fmt.Println("Processing index file: " + path)
			targetFile := path
			pointToFile := strings.Replace(targetFile, markdownPath, webPath, 1)
			pointToFile = strings.Replace(pointToFile, "md", "html", 1)
			mdFile := markdown{
				targetFile:  targetFile,
				pointToFile: pointToFile,
			}
			setIndexHtml(mdFile, webPath)

		}
		return nil
	})
	fmt.Println("build complete")
	if err != nil {
		log.Fatal(err)
	}

}
