package noc

import "net/http"

type Routable interface {
	Route(method string, path string, handler HandlerFunc)
}

type Router interface {
	Routable
	http.Handler
}
