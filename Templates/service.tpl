package service

{{MakeImports .Imports}}
{{$intface:=.Name}}
//{{.Name}}接口的实现类
type {{MakeServiceName $intface}} struct{

}

func New{{MakeServiceName $intface}} () *{{MakeServiceName $intface}} {
    return &{{MakeServiceName $intface}}{}
}

{{range .MethodsNew}}
func (this *{{MakeServiceName $intface}}) {{.FunName}}({{MakeParams .FunParams}}) {{MakeResults .FunResult}} {
    // your codes
}

{{end}}