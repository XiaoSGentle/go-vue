package xtemplate

import (
	"log"
	"os"
	"text/template"
)

func GetTemplate(tmplPath string, funcMap template.FuncMap) *template.Template {
	file, err := os.ReadFile(tmplPath)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	// 模板解析
	return template.Must(template.New(string(file)).Funcs(funcMap).Parse(string(file)))
}
