package Commands

import (
	"flag"
	"github.com/hhaojin/generate/Helper"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func init() {

	vc := NewResourceCmd()
	vc.Register("ResourceCommand", vc) //注册命令

}

type ResourceCommand struct { //用来处理静态资源的生成
	IsCreate *bool
	ResPath  *string
	CommandSet
	ServiceCommandSet *flag.FlagSet //这里做了改动。直接在属性中
}

func NewResourceCmd() *ResourceCommand {
	is_create := false
	return &ResourceCommand{IsCreate: &is_create,
		ServiceCommandSet: flag.NewFlagSet("resource args", flag.ExitOnError)}
}
func (this *ResourceCommand) Init() {
	//go run main.go resource
	if len(os.Args) > 2 && os.Args[1] == "resource" {
		this.IsCreate = this.ServiceCommandSet.Bool("c", false, "create resource")
		this.ResPath = this.ServiceCommandSet.String("p", "/Templates", "resource path (default:./Templates")
		err := this.ServiceCommandSet.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)
		}
	}
}
func (this *ResourceCommand) Run() {
	if *this.IsCreate { //-c 有
		dir := Helper.WorkDir + *this.ResPath
		//遍历文件
		res := Helper.LoadResource(dir) //加载 资源文件
		tpl, err := ioutil.ReadFile(Helper.WorkDir + "/resource/resource.tpl")
		if err != nil {
			log.Fatal("resource.tpl error")
		}
		if res != nil {
			tmpl, err := template.New("resource").Funcs(Helper.NewTplFunction()).Parse(string(tpl))
			if err != nil {
				log.Fatal("resource parse error:", err)
			}
			file, err := os.OpenFile(Helper.WorkDir+"/resource/static.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				log.Fatal("load resource  error:", err)
			}
			err = tmpl.Execute(file, res)
			if err != nil {
				log.Fatal("create resource  error:", err)
			}
			log.Println("资源文件刷新成功")
		}
	}
}
