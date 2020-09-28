package models
{{$ModelName:=(printf "%s" .TableName) | CamelCase}}
type {{$ModelName}} struct {
    {{range .Data}}{{$fieldName:=(index . "Field")}}{{index . "Field" | CamelCase}} {{index . "Type" | MapType}} `gorm:"cloumn:{{$fieldName}}" json:"{{$fieldName}}" form:"{{$fieldName}}"`
    {{end}}
}
{{range .Data}}
func (this *{{$ModelName}}) With{{index . "Field" | CamelCase}}({{index . "Field" }} {{index . "Type" | MapType}}) *{{$ModelName}} {
    this.{{index . "Field" | CamelCase}} = {{index . "Field"}}
    return this
}
{{end}}

{{$ModelOptions:=(printf "%sOptions" $ModelName)}}
type {{$ModelOptions}} struct {
    apply func(*{{$ModelName}})
}

func New{{$ModelName}}(opts... *{{$ModelName}}Options) *{{$ModelName}} {
    m:= &{{$ModelName}}{}
    for _,opt:=range opts{
       opt.apply(m)
    }
    return m
}
{{range .Data}}
func {{$ModelName}}{{index . "Field" | CamelCase}}({{index . "Field"}} {{index . "Type" | MapType}}) *{{$ModelOptions}} {
    return &{{$ModelOptions}}{func(model *{{$ModelName}}) {
        model.{{index . "Field" | CamelCase}} = {{index . "Field"}}
    }}
}
{{end}}









