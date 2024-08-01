package noc

import "net/http"

type Router interface {
	Route(method string, path string, handler HandlerFunc)
	http.Handler
}
