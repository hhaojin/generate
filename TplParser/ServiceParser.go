package TplParser

import (
	"github.com/hhaojin/generate/AstParser"
	"github.com/hhaojin/generate/Helper"
	"github.com/hhaojin/generate/resource"
	"log"
	"os"
	"text/template"
)

type ServiceParser struct {
	TplContent string //模板内容
}

//初始化 并读取模板内容
func NewServiceParser() *ServiceParser {
	return &ServiceParser{TplContent: Helper.UnGzip(resource.SERVICE_TPL)}
}

//直接生成到 目标文件中,target 只需要传文件名即可
func (this *ServiceParser) Parse(pi *AstParser.PInterface, dir string, target string) {
	tpl := template.New("service").Funcs(Helper.NewTplFunction())

	tmpl, err := tpl.Parse(this.TplContent)
	if err != nil {
		log.Fatal("service tpl parse-err:", err)
	}
	openFile, err := Helper.CreateMutiDir(dir, target)
	if err != nil {
		log.Fatal("mkdir error:", err)
	}
	file, err := os.OpenFile(openFile,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("service target error:", err)
	}

	err = tmpl.Execute(file, pi)
	if err != nil {
		log.Fatal("gen service error:", err)
	}

}
