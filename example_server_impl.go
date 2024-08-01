package noc

import (
	"log"
	"net/http"
)

var _ Server = &ExampleServer{}

type ExampleServer struct {
	ServerName string
	router     Router
}

func (s *ExampleServer) Route(method string, path string, handlerFunc HandlerFunc) {
	s.router.Route(method, path, handlerFunc)
}

func (s *ExampleServer) GET(path string, handler HandlerFunc) {
	s.Route("GET", path, handler)
}

func (s *ExampleServer) POST(path string, handler HandlerFunc) {
	s.Route("POST", path, handler)
}

func (s *ExampleServer) Start(address string) error {
	log.Printf("Server %s is listening on %s", s.ServerName, address)
	return http.ListenAndServe(address, s.router)
}

func NewExampleServer(serverName string, router Router) *ExampleServer {
	return &ExampleServer{
		ServerName: serverName,
		router:     router,
	}
}
