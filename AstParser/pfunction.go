package AstParser

import (
	"go/ast"
)

type PFunction struct {
	FunName   string    //函数名称 ...  没实现部分: 返回值
	FunParams []*PParam //参数
	FunResult []*PParam //返回值集合
}

func NewPFunction() *PFunction {
	return &PFunction{FunParams: make([]*PParam, 0), FunResult: make([]*PParam, 0)}
}

//加入参数
func (this *PFunction) AddParam(pp *PParam) {
	this.FunParams = append(this.FunParams, pp)
}

//加入返回值
func (this *PFunction) AddResult(rr *PParam) {
	this.FunResult = append(this.FunResult, rr)
}
func (this *PFunction) Visit(node ast.Node) ast.Visitor {
	if f, ok := node.(*ast.Field); ok {
		if fn, ok := f.Type.(*ast.FuncType); ok {
			this.FunName = f.Names[0].Name     //取到函数名
			for _, p := range fn.Params.List { //处理参数
				pname := ""
				if len(p.Names) > 0 {
					pname = p.Names[0].Name
				}
				pp := NewPParam(pname, p.Type) //获取参数名的时候 需要判断是否有参数名
				this.AddParam(pp)
			}
			if fn.Results != nil { //这里要加一步判断，因为返回值可能没有
				for _, r := range fn.Results.List { //处理返回值
					rname := ""
					if len(r.Names) > 0 {
						rname = r.Names[0].Name
					}
					rr := NewPParam(rname, r.Type)
					this.AddResult(rr)
				}
			}

		}
	}
	return this
}

//判断对象是否是Function ,其实就是先找*ast.Field，然后判断type是否是*ast.FunType
func Function(decl *ast.Field) *PFunction {
	pf := NewPFunction()
	ast.Walk(pf, decl)
	if pf.FunName != "" {
		return pf
	}
	return nil
}

//判断对象是否是Function切片 。是上个函数的复数形式
func FunctionList(flist *ast.FieldList) []*PFunction {
	ret := make([]*PFunction, 0)
	if flist == nil || flist.List == nil || len(flist.List) == 0 {
		return ret
	}
	for _, field := range flist.List {
		if f := Function(field); f != nil {
			ret = append(ret, f)
		}
	}
	return ret
}
