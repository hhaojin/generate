package gin_

type ServiceBuilder struct {
	service      interface{}
	endpoint     Endpoint
	requestFunc  EncodeRequestFunc
	responseFunc DecodeResponseFunc
	middlewares  []Middleware
}

func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{}
}

func (this *ServiceBuilder) Service(service interface{}) *ServiceBuilder {
	this.service = service
	return this
}

func (this *ServiceBuilder) Endpoint(endpoint Endpoint) *ServiceBuilder {
	this.endpoint = endpoint
	return this
}

func (this *ServiceBuilder) RequestFunc(requestFunc EncodeRequestFunc) *ServiceBuilder {
	this.requestFunc = requestFunc
	return this
}

func (this *ServiceBuilder) ResponseFunc(responseFunc DecodeResponseFunc) *ServiceBuilder {
	this.responseFunc = responseFunc
	return this
}

func (this *ServiceBuilder) Middleware(middleware Middleware) *ServiceBuilder {
	this.middlewares = append(this.middlewares, middleware)
	return this
}

func (this *ServiceBuilder) BuildServer(path string, method string) {

	finalEndpoint := this.endpoint
	for _, middle := range this.middlewares {
		finalEndpoint = middle(finalEndpoint)
	}

	handler := RegisterHandler(finalEndpoint, this.requestFunc, this.responseFunc)
	SetServerHandler(method, path, handler)
}
