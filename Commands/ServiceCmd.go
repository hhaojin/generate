package Commands

import (
	"flag"
	"fmt"
	"github.com/hhaojin/generate/AstParser"
	"github.com/hhaojin/generate/Helper"
	"github.com/hhaojin/generate/TplParser"
	"log"
	"os"
)

func init() {
	vc := NewServiceCmd()
	vc.Register("ServiceCommand", vc) //注册命令
}

type ServiceCommand struct { //用来生成 服务相关的工具 命令，基于interface
	IsCreate *bool   //是否是创建代码  。也许后面有 rm 、update
	IFile    *string //接口文件
	Dir      *string //目标文件夹
	CommandSet
	ServiceCommandSet *flag.FlagSet //这里做了改动。直接在属性中
}

func NewServiceCmd() *ServiceCommand {
	is_create := false
	return &ServiceCommand{IsCreate: &is_create, //这里做了改动，由于IsCreate是指针，需要赋值，否则后面判断会出错
		ServiceCommandSet: flag.NewFlagSet("service args", flag.ExitOnError)}
}
func (this *ServiceCommand) Init() {
	//go run main.go service -c
	if len(os.Args) > 2 && os.Args[1] == "service" {
		this.IsCreate = this.ServiceCommandSet.Bool("c", false, "create service")
		this.IFile = this.ServiceCommandSet.String("i", "", "interface file (app\\api)")
		this.Dir = this.ServiceCommandSet.String("d", "app/service",
			"generates to the specified folder (default:app\\service)")
		err := this.ServiceCommandSet.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)
		}
	}
}
func (this *ServiceCommand) Run() {
	if *this.IsCreate { //-c 有
		if Helper.IsApiFileExist(*this.IFile) {
			p := AstParser.NewParser(Helper.GetApiFile(*this.IFile))
			infs := p.ParseInterfaces()                  //切片
			tplSevParser := TplParser.NewServiceParser() //模板解析类，专门处理service 解析和生成
			for _, pi := range infs {                    //遍历接口切片
				tplSevParser.Parse(pi, *this.Dir, Helper.MakeServiceName(pi.Name)) //直接生成了文件
			}
			fmt.Println("创建成功")
		} else {
			log.Println("interface file error")
		}
	}
}
