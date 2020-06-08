package AstParser

import (
	"go/ast"
	"sort"
)

type FieldMap map[int]string

//排序 取出key 默认从小到大
func (this FieldMap) Keys() []int {
	keys := make([]int, 0)
	for key := range this {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

//根据排序过后的key取出value集合
func (this FieldMap) Values() []interface{} {
	ret := make([]interface{}, 0)
	keys := this.Keys()
	for _, key := range keys {
		ret = append(ret, this[key])
	}
	return ret
}

type PParam struct {
	Name     string
	FieldPos FieldMap //记录每个exp的起始pos
}

func NewPParam(name string, t ast.Expr) *PParam {
	ret := &PParam{name,
		make(map[int]string, 0)}
	ast.Walk(ret, t)
	return ret
}

func (this *PParam) Visit(node ast.Node) ast.Visitor {
	switch exp := node.(type) {
	case *ast.Ident: //这一步很重要。各自的类型 经过递归后，大部分都会回归到这个类型
		this.FieldPos[int(exp.Pos())] = exp.Name //记录 起始位置
	case *ast.ArrayType: //切片类型
		this.FieldPos[int(exp.Pos())] = "[]"
	case *ast.InterfaceType: //interface类型
		this.FieldPos[int(exp.Pos())] = "interface{}"
	case *ast.MapType: //map 类型

		this.FieldPos[int(exp.Map)] = "map"
		this.FieldPos[int(exp.Key.Pos())-1] = "["
		this.FieldPos[int(exp.Key.End())] = "]"
	case *ast.SelectorExpr:
		this.FieldPos[int(exp.X.End())] = "."
		break
	case *ast.StarExpr: //指针
		this.FieldPos[int(exp.Pos())] = "*"
	case *ast.ChanType: //chan类型
		this.FieldPos[int(exp.Pos())] = "chan " //空格必须要
	}
	return this
}
