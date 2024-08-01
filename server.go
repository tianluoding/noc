package noc

type HandlerFunc func(ctx *Context)

type Server interface {
	Route(method string, path string, handler HandlerFunc)
	Start(address string) error
}
