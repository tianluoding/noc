package noc

import (
	"net/http"

	"github.com/tianluoding/noc/logger"
)

var _ Server = &ExampleServer{}

type ExampleServer struct {
	ServerName string
	router     Router
	filters    []FilterFunc
}

func (s *ExampleServer) Route(method string, path string, handlerFunc HandlerFunc, filters ...FilterFunc) {
	// deep copy
	copedFilters := make([]FilterFunc, 0, len(s.filters)+len(filters)+2)
	copy(copedFilters, s.filters)
	copedFilters = append(copedFilters, filters...)
	s.router.Route(method, path, handlerFunc, copedFilters...)
}

func (s *ExampleServer) GET(path string, handler HandlerFunc) {
	s.Route("GET", path, handler, s.filters...)
}

func (s *ExampleServer) POST(path string, handler HandlerFunc) {
	s.Route("POST", path, handler, s.filters...)
}

func (s *ExampleServer) Start(address string) error {
	logger.Logger.Infof("Server %s is listening on %s", s.ServerName, address)
	return http.ListenAndServe(address, s.router)
}

func (s *ExampleServer) AddFilters(filters ...FilterFunc) {
	s.filters = append(s.filters, filters...)
}

func NewExampleServer(serverName string) *ExampleServer {
	router := NewMapRouter()

	return &ExampleServer{
		ServerName: serverName,
		router:     router,
		filters:    make([]FilterFunc, 0, 4),
	}
}
