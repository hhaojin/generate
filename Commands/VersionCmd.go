package Commands

import (
	"flag"
	"fmt"
)

type VersionCmd struct {
	CommandSet
	FlagV *bool
}

func init() {
	vc := NewVersionCmd()
	vc.Register("VersionCmd", vc)
}

func NewVersionCmd() *VersionCmd {
	return &VersionCmd{}
}

func (this *VersionCmd) Init() {
	//go run main.go -v
	this.FlagV = flag.Bool("v", false, "show version")
}

func (this *VersionCmd) Run() {
	if *this.FlagV {
		fmt.Printf("version: %s", "1.0")
	}
}
