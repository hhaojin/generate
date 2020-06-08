package gin_

import "github.com/gin-gonic/gin"

var ServiceHandlerMap map[string]map[string]gin.HandlerFunc

func init() {
	ServiceHandlerMap = make(map[string]map[string]gin.HandlerFunc)
}

func SetServerHandler(method string, path string, handler gin.HandlerFunc) {
	if _, ok := ServiceHandlerMap[path]; !ok {
		ServiceHandlerMap[path] = make(map[string]gin.HandlerFunc)
	}
	ServiceHandlerMap[path][method] = handler
}

func BootStrap(router *gin.Engine) {
	for path, hmap := range ServiceHandlerMap {
		for method, handler := range hmap {
			router.Handle(method, path, handler)
		}
	}
}
