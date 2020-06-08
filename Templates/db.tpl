package models
{{$ModelName:=(printf "%s%s" .TableName "Model") | Ucfirst}}
type {{$ModelName}} struct {
    {{range .Data}}{{index . "Field" | Ucfirst}} {{index . "Type" | MapType}} `db:"{{index . "Field"}}"`
    {{end}}
}
{{range .Data}}
func(this *{{$ModelName}}) With{{index . "Field" | Ucfirst}}({{index . "Field" }} {{index . "Type" | MapType}}) *{{$ModelName}} {
    this.{{index . "Field" | Ucfirst}}={{index . "Field"}}
    return this
}
{{end}}

{{$ModelOptions:=(printf "%sOptions" $ModelName)}}
type {{$ModelOptions}} struct {
    apply func(*{{$ModelName}})
}

func New{{$ModelName}}(opts...*{{$ModelName}}Options) *{{$ModelName}} {
    m:= &{{$ModelName}}{}
    for _,opt:=range opts{
       opt.apply(m)
    }
    return m
}
{{range .Data}}
func {{$ModelName}}{{index . "Field" | Ucfirst}}({{index . "Field"}} {{index . "Type" | MapType}}) *{{$ModelOptions}} {
    return &{{$ModelOptions}}{func(model *{{$ModelName}}) {
        model.{{index . "Field" | Ucfirst}}={{index . "Field"}}
    }}
}
{{end}}









