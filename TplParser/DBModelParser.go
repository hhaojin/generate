package TplParser

import (
	"gitee.com/hhaojin/generate/Helper"
	"gitee.com/hhaojin/generate/resource"
	"log"
	"os"
	"text/template"
)

type DBModelParser struct {
	TplContent string //模板内容
}

//初始化 并读取模板内容
func NewDBModelParser() *DBModelParser {
	return &DBModelParser{TplContent: Helper.UnGzip(resource.DB_TPL)}
}
func (this *DBModelParser) Parse(data interface{}, dir string, target string) {
	tpl := template.New("DBModel").Funcs(Helper.NewTplFunction())
	tmpl, err := tpl.Parse(this.TplContent)
	if err != nil {
		log.Fatal("db tpl parse-error:", err)
	}
	openFile, err := Helper.CreateMutiDir(dir, target)
	if err != nil {
		log.Fatal("mkdir error:", err)
	}

	file, err := os.OpenFile(openFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("model target error:", err)
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		log.Fatal("generate model error:", err)
	}
}
