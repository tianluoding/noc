package noc

import "net/http"

type Context struct {
	R *http.Request
	W http.ResponseWriter
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{R: r, W: w}
}
