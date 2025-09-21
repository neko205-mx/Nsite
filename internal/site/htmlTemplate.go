package site

import (
	"log"
	"os"
)

// LoadTemplate 加载指定的模板文件
func LoadTemplate(templatePath string, fileName string) []byte {
	load, err := os.ReadFile(templatePath + fileName)
	if err != nil {
		log.Fatal(err)
	}
	return load
}
