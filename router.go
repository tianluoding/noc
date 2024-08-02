package noc

import "net/http"

type Routable interface {
	Route(method string, path string, handler HandlerFunc, filters ...FilterFunc)
}

type Router interface {
	Routable
	http.Handler
}
