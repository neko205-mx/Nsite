package site

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// fileCopy 将文件从markroot目录复制到wwwroot目录对应位置
// filename参数指定源文件的完整路径
// 如果目标目录不存在会自动创建
func fileCopy(filename string) {
	pointToFile := strings.Replace(filename, markroot, wwwroot, 1)

	//检查目标目录
	pointDir := filepath.Dir(pointToFile)
	if err := os.MkdirAll(pointDir, 0755); err != nil {
		log.Fatal(err)
	}

	//创建目标文件
	destFile, err := os.OpenFile(pointToFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//打开文件写入
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(destFile, file)

}

// fileProcessing 将指定的Markdown文件转换为HTML文件
// 读取markdownFile.targetFile指定的Markdown文件内容，使用goldmark转换为HTML格式
// 将转换后的HTML内容写入到markdownFile.pointToFile指定的路径
func fileProcessing(markdownFile markdown) {
	fmt.Println(markdownFile.pointToFile)
	fmt.Println(markdownFile.targetFile)

	var buf bytes.Buffer
	var data, _ = os.ReadFile(markdownFile.targetFile)

	if err := MdFull().Convert(data, &buf); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

	wwwDir := filepath.Dir(markdownFile.pointToFile)
	err := os.MkdirAll(wwwDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	htmlFile, err := os.Create(markdownFile.pointToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer htmlFile.Close()

	_, err = htmlFile.Write([]byte(pagehtmlHeader))
	if err != nil {
		return
	}
	// 写入html文件
	_, err = htmlFile.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	// 写入结尾
	_, _ = htmlFile.Write([]byte(htmlend))
}

// setIndexHtml 处理markdown文件并生成对应的HTML索引文件
// 1. 读取目标markdown文件作为模板
// 2. 遍历wwwroot目录下的所有HTML文件，生成markdown格式的链接并追加到模板内容
// 3. 将处理后的内容写入目标markdown文件
// 4. 使用goldmark将markdown转换为HTML并写入指定文件
// 5. 清理模板文件，恢复原始内容
// 参数markdownFile包含目标文件路径(targetFile)和输出HTML文件路径(pointToFile)
func setIndexHtml(markdownFile markdown) {
	//将index作为模板读取
	indexTemplate, _ := os.ReadFile(markdownFile.targetFile)
	indexTemplateProcessing := indexTemplate
	fmt.Println("index:" + markdownFile.targetFile) //index:wwwmark/index.md
	err := filepath.WalkDir(wwwroot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("indexHtml:" + path)

		if strings.HasSuffix(path, ".html") {
			urlPath := strings.Replace(path, wwwroot, ".", 1)
			fmt.Println(urlPath)
			urlName := filepath.Base(urlPath)
			urlName = strings.TrimSuffix(urlName, ".html")

			mdUrl := fmt.Sprintf(" - [%s](%s)\n\n", urlName, urlPath)
			fmt.Println(mdUrl)

			// 不直接追加写入
			indexTemplateProcessing = append(indexTemplateProcessing, []byte(mdUrl)...)
		}
		return nil
	})
	//只写 创建 覆盖
	indexMdFile, err := os.OpenFile(markdownFile.targetFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer indexMdFile.Close()
	_, err = indexMdFile.Write([]byte(indexTemplateProcessing))
	if err != nil {
		log.Fatal(err)
	}

	println(markdownFile.targetFile)

	var buf bytes.Buffer
	var data, _ = os.ReadFile(markdownFile.targetFile)

	if err := MdFull().Convert(data, &buf); err != nil {
		log.Fatal(err)
	}
	htmlFile, err := os.Create(markdownFile.pointToFile)
	if err != nil {
		log.Fatal(err)
	}
	//添加头
	defer htmlFile.Close()
	_, err = htmlFile.Write([]byte(indexhtmlHeader))
	if err != nil {
		return
	}
	//写入index.html
	_, err = htmlFile.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	//添加尾
	_, _ = htmlFile.Write([]byte(htmlend))

	// 清理index md模板
	cleanIndexTemplateData, err := os.OpenFile(markdownFile.targetFile, os.O_WRONLY|os.O_TRUNC, 0775)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanIndexTemplateData.Close()
	_, err = cleanIndexTemplateData.Write(indexTemplate)

}

// Clean 删除wwwroot目录及其所有内容
func Clean() {
	err := os.RemoveAll(wwwroot)
	if err != nil {
		return
	}
}
