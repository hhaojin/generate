package TplParser

import (
	"github.com/hhaojin/generate/Helper"
	"github.com/hhaojin/generate/resource"
	"log"
	"os"
	"text/template"
)

type LibParser struct {
	TplContent string //模板内容
}

//初始化 并读取模板内容
func NewLibParser() *LibParser {
	return &LibParser{TplContent: Helper.UnGzip(resource.LIB_TPL)}
}
func (this *LibParser) Parse(data interface{}, openfile string) {
	tpl := template.New("LibTpl").Funcs(Helper.NewTplFunction())
	tmpl, err := tpl.Parse(this.TplContent)
	if err != nil {
		log.Fatal("lib tpl parse-error:", err)
	}
	file, err := os.OpenFile(openfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("lib target error:", err)
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		log.Fatal("generate lib error:", err)
	}
	log.Printf("生成成功,请在%s下查看文件", openfile)
}
