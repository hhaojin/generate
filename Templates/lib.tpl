package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/hhaojin/generate/framework/gin_"
)

func init() {
    {{$serviceName := printf "%sService" .InterfaceName}}
	{{$serviceName}} := New{{.InterfaceName}}Impl()
	gin_.NewServiceBuilder().Service({{$serviceName}}).
		Middleware({{.InterfaceName}}_Middleware()).
		Endpoint({{.InterfaceName}}_Endpoint({{$serviceName}})).
		RequestFunc({{.InterfaceName}}_Request()).
		ResponseFunc({{.InterfaceName}}_Response()).
		BuildServer("{{.RoutePath}}","{{.MethodName}}")
}

func {{.InterfaceName}}_Middleware() gin_.Middleware {
	return func(next gin_.Endpoint) gin_.Endpoint {
		return func(context *gin.Context, request interface{}) (response interface{}, err error) {
			return next(context, request)
		}
	}
}

func {{.InterfaceName}}_Endpoint(test *{{.InterfaceName}}) gin_.Endpoint {
	return func(context *gin.Context, request interface{}) (response interface{}, err error) {
		//请在这里写 具体的业务调用，就是service调用
		return nil, nil
	}
}

func {{.InterfaceName}}_Request() gin_.EncodeRequestFunc {
	return func(context *gin.Context) (i interface{}, e error) {
		 //请在这里写 如何绑定参数
		return nil, nil
	}
}

func {{.InterfaceName}}_Response() gin_.DecodeResponseFunc {
	return func(context *gin.Context, response interface{}) error {
		context.JSON(200, response)
		return nil
	}
}