package AstParser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Parser struct {
	asf *ast.File //ast对象
}

//所谓的构造函数
func NewParser(fname string) *Parser {
	fset := token.NewFileSet()
	ast_file, err := parser.ParseFile(fset, fname, nil, 0|parser.ParseComments)
	if err != nil {
		panic(fmt.Sprintf("parser init err:%s", err.Error()))
	}
	return &Parser{asf: ast_file}
}

//把代码的 接口对象全部取出来
func (this *Parser) ParseInterfaces() []*PInterface {
	ret := make([]*PInterface, 0)

	for _, dec := range this.asf.Decls { //这里处理的是定义（如接口、func 、var等等)
		if i := Interface(dec); i != nil { //判断是否是接口
			i.Imports = this.asf.Imports
			ret = append(ret, i)
		}

	}
	return ret
}
