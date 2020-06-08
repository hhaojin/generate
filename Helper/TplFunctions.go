package Helper

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"gitee.com/hhaojin/generate/AstParser"
	"go/ast"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

//给模板使用的函数集
func NewTplFunction() template.FuncMap {
	fm := make(map[string]interface{})
	fm["CamelCase"] = CamelCase
	fm["SnakeCase"] = SnakeCase
	fm["Ucfirst"] = Ucfirst
	fm["MakeServiceName"] = MakeServiceName
	fm["MakeParams"] = MakeParams
	fm["MakeImports"] = MakeImports
	fm["MakeResults"] = MakeResults
	fm["Gzip"] = Gzip
	fm["MapType"] = MapDBType
	return fm
}

//生成Service名称。 一般规则是接口名+Impl
func MakeServiceName(sname string) string {
	return CamelCase(sname) + "Impl"
}

//解析函数参数
func MakeParams(pps []*AstParser.PParam) string {
	ret := ""
	for index, p := range pps {
		if index != 0 {
			ret += ","
		}
		indices := p.FieldPos.Keys() //排序
		format_str := ""
		for _ = range indices {
			format_str += "%s"
		}
		vs := []interface{}{p.Name}
		vs = append(vs, p.FieldPos.Values()...)
		format_str = strings.Trim(format_str, ".")
		ret += fmt.Sprintf("%s "+format_str, vs...)
	}

	return strings.TrimLeft(ret, " ")

}

//解析返回值参数
func MakeResults(pps []*AstParser.PParam) string {
	ret := MakeParams(pps)
	lens := strings.Split(ret, " ")
	if len(lens) > 1 {
		return "(" + ret + ")"
	}
	return ret
}

//解析imports
func MakeImports(impts []*ast.ImportSpec) string {
	if len(impts) == 0 {
		return ""
	}
	name := func(n *ast.Ident) string {
		if n != nil {
			return n.Name
		}
		return ""
	}

	if len(impts) == 1 {
		return fmt.Sprintf("import %s%s", name(impts[0].Name), impts[0].Path.Value)
	}
	ret := "import (\n"
	for _, impt := range impts {
		ret += fmt.Sprintf(" %s%s\n", name(impt.Name), impt.Path.Value)
	}
	ret += ")\n"
	return ret
}

func Gzip(str string) string {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	_, err := gz.Write([]byte(str))
	if err != nil {
		log.Println(err)
		return ""
	}
	err = gz.Close() //这里要关掉，否则取不到数据  也可手工flush.但依然要关掉gz
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
func UnGzip(str string) string {
	dbytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println(err)
		return ""
	}
	read_data := bytes.NewReader(dbytes)
	reader, err := gzip.NewReader(read_data)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer reader.Close()
	ret, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println("read gzip error:", err)
		return ""
	}

	return string(ret)

}
