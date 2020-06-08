package Commands

import (
	"flag"
	"fmt"
	"gitee.com/hhaojin/generate/Helper"
	"gitee.com/hhaojin/generate/TplParser"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	vc := NewLibGenCmd()
	vc.Register("LibGenerateCommand", vc) //注册命令

}

type LibEntity struct {
	InterfaceName string
	RoutePath     string
	MethodName    string
}

func NewLibEntity(s1 string, s2 string, s3 string) *LibEntity {
	return &LibEntity{s1, s2, s3}
}
func setLibEntity(args ...string) *LibEntity {
	if len(args) != 3 {
		return nil
	}
	return NewLibEntity(args[0], args[1], args[2])
}

type LibGenCmd struct { //用来生成三层架构代码模板
	Dest          *string
	InterfaceName *string
	RoutePath     *string
	MethodName    *string
	CommandSet
	ServiceCommandSet *flag.FlagSet
}

func NewLibGenCmd() *LibGenCmd {
	dest := ""
	return &LibGenCmd{Dest: &dest, ServiceCommandSet: flag.NewFlagSet("LibGenerate args", flag.ExitOnError)}
}
func (this *LibGenCmd) Init() {
	if len(os.Args) > 2 && os.Args[1] == "lib" {
		this.Dest = this.ServiceCommandSet.String("c", "app/lib", "destination (for example:/app/lib)")
		this.InterfaceName = this.ServiceCommandSet.String("i", "", "your interface name")
		this.RoutePath = this.ServiceCommandSet.String("r", "", "route path")
		this.MethodName = this.ServiceCommandSet.String("m", "GET", "method(GET or POST)")
		err := this.ServiceCommandSet.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)
		}
	}
}
func (this *LibGenCmd) Run() {
	if strings.Trim(*this.Dest, " ") == "" {
		return
	}
	if strings.Trim(*this.InterfaceName, " ") == "" {
		log.Fatal("param -i empty")
	}
	if strings.Trim(*this.RoutePath, " ") == "" {
		log.Fatal("param -r empty")
	}

	//items := []string{" your interface name", " route path", " method(GET or POST)"}
	//replies := []string{}
	//in := bufio.NewScanner(os.Stdin)
	//for _, item := range items {
	//	fmt.Println(item)
	//	for in.Scan() {
	//		str := in.Text()
	//		if strings.Trim(str, " ") == "" {
	//			fmt.Println(item)
	//			continue
	//		} else {
	//			replies = append(replies, str)
	//			break
	//		}
	//	}
	//}

	e := setLibEntity(*this.InterfaceName, *this.RoutePath, *this.MethodName)
	target := fmt.Sprintf("%sLib_%d", e.InterfaceName, time.Now().Unix())
	openFile, err := Helper.CreateMutiDir(*this.Dest, target)
	if err != nil {
		log.Fatal("dir error:", err)
	}
	if e == nil {
		log.Fatal("error lib params")
	}
	libParser := TplParser.NewLibParser()
	libParser.Parse(e, openFile)
}
