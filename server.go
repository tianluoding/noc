package noc

type HandlerFunc func(ctx *Context) error

type Server interface {
	Routable
	Start(address string) error
	AddFilters(filters ...FilterFunc)
}
