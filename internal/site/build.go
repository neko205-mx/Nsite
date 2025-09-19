package site

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type markdown struct {
	targetFile  string
	pointToFile string
}

var markroot = "wwwmark"
var wwwroot = "wwwroot"

// Build 遍历 markroot 目录下的文件，处理所有 .css 文件和包含 "page" 的 .md 文件，
// 将它们复制或转换为对应的 .html 文件到 wwwroot 目录下。
// 同时查找并处理 index.md 文件作为网站的首页。
func Build() {
	Clean()
	err := filepath.WalkDir(markroot, func(path string, d os.DirEntry, err error) error {
		if strings.Contains(path, ".css") {
			fileCopy(path)
		}

		if err != nil {
			return err
		}
		if strings.Contains(path, "page") && strings.Contains(path, "md") {
			//fmt.Println(path)
			targetFile := path
			//fmt.Println(targetFile)
			pointToFile := strings.Replace(targetFile, markroot, wwwroot, 1)
			pointToFile = strings.Replace(pointToFile, "md", "html", 1)
			//fmt.Println(pointToFile)

			mdFile := markdown{
				targetFile:  targetFile,
				pointToFile: pointToFile,
			}
			fileProcessing(mdFile)
		}
		//fmt.Printf("%s\n", path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(markroot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, "index.md") {
			fmt.Println(path)
			targetFile := path
			pointToFile := strings.Replace(targetFile, markroot, wwwroot, 1)
			pointToFile = strings.Replace(pointToFile, "md", "html", 1)
			mdFile := markdown{
				targetFile:  targetFile,
				pointToFile: pointToFile,
			}
			fmt.Println(mdFile)
			setIndexHtml(mdFile)

		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}
