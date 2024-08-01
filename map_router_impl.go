package noc

import "net/http"

var _ Router = &MapRouter{}

type MapRouter struct {
	handlers map[string]HandlerFunc
}

func (h *MapRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		ctx := NewContext(w, r)
		if err := handler(ctx); err != nil {
			ctx.BadRequestJSON()
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *MapRouter) Route(method string, path string, handler HandlerFunc) {
	h.handlers[h.key(method, path)] = handler
}

func (h *MapRouter) key(method string, path string) string {
	return method + "#" + path
}

func NewMapRouter() *MapRouter {
	return &MapRouter{
		handlers: make(map[string]HandlerFunc),
	}
}
