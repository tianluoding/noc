package noc

type HandlerFunc func(ctx *Context) error

type Server interface {
	Route(method string, path string, handler HandlerFunc)
	Start(address string) error
}
