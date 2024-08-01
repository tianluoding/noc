package noc

import "net/http"

type ExampleServer struct {
	ServerName string
}

func (s *ExampleServer) Route(method string, path string, handler HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(w, r)
		handler(ctx)
	})
}

func (s *ExampleServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewExampleServer(serverName string) *ExampleServer {
	return &ExampleServer{
		ServerName: serverName,
	}
}
