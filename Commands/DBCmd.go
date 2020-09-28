package Commands

import (
	"flag"
	"fmt"
	"github.com/hhaojin/generate/Helper"
	"github.com/hhaojin/generate/TplParser"
	"log"
	"os"
)

func init() {
	vc := NewDBCmd()
	vc.Register("DBCommand", vc) //注册命令
}

type DBCommand struct { //用来处理静态资源的生成
	DSN       *string
	TableName *string
	Dir       *string
	CommandSet
	ServiceCommandSet *flag.FlagSet
}

func NewDBCmd() *DBCommand {
	var dsn, tablename string = "", ""
	return &DBCommand{DSN: &dsn, TableName: &tablename,
		ServiceCommandSet: flag.NewFlagSet("db args", flag.ExitOnError)}
}
func (this *DBCommand) Init() {
	//go run main.go resource
	if len(os.Args) > 2 && os.Args[1] == "db" {
		this.TableName = this.ServiceCommandSet.String("m", "", "create model")
		this.Dir = this.ServiceCommandSet.String("d", "app/models",
			"generates to the specified folder (default:app\\models)")
		err := this.ServiceCommandSet.Parse(os.Args[2:])
		if err != nil {
			log.Println(err)
		}
	}
}

func (this *DBCommand) Run() {
	config, err := Helper.LoadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("missing config-file: .%s", Helper.SYS_CONFIG_PATH))
	}
	if *this.TableName != "" { // 生成模型
		db := Helper.NewDB(config.DB.Driver, config.DB.DSN)
		data, err := db.DescTable(*this.TableName, config.DB.Prefix) //DBModel 包含了 tablename 以及 字段名等
		if err != nil {
			log.Fatal("model error:", err.Error())
		}
		dm := TplParser.NewDBModelParser()

		fmt.Println(Helper.CamelCase(data.TableName))
		dm.Parse(data, *this.Dir, Helper.CamelCase(data.TableName))
		fmt.Println("生成模型")
	}
}
