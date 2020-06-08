package AstParser

import (
	"go/ast"
)

type PInterface struct {
	Name       string //接口名称
	Methods    *ast.FieldList
	MethodsNew []*PFunction
	Imports    []*ast.ImportSpec //import的部分
}

func NewPInterface() *PInterface {
	return &PInterface{}
}
func (this *PInterface) Visit(node ast.Node) ast.Visitor {
	if ts, ok := node.(*ast.TypeSpec); ok {
		if t, ok := ts.Type.(*ast.InterfaceType); ok {
			this.Name = ts.Name.Name
			this.Methods = t.Methods
			this.MethodsNew = FunctionList(t.Methods) //解析方法 映射成我们自己的对象:PFunction
		}
	}
	return this
}

//用来判断 当前类型是否是interface接口
func Interface(decl ast.Decl) *PInterface {
	ret := NewPInterface()
	ast.Walk(ret, decl)
	if ret.Methods != nil {
		return ret
	}
	return nil

}
