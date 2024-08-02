package noc

import (
	"fmt"
	"net/http"

	"github.com/tianluoding/noc/logger"
)

type HandlerFunc func(ctx *Context) error

type Server interface {
	Routable
	Start(address string) error
	AddFilters(filters ...FilterFunc)
}

var _ Server = &Noc{}

type Noc struct {
	ServerName string
	router     Router
	filters    []FilterFunc
}

func (s *Noc) Route(method string, path string, handlerFunc HandlerFunc, filters ...FilterFunc) {
	// deep copy
	copedFilters := make([]FilterFunc, 0, len(s.filters)+len(filters)+2)
	copy(copedFilters, s.filters)
	copedFilters = append(copedFilters, filters...)
	s.router.Route(method, path, handlerFunc, copedFilters...)
}

func (s *Noc) GET(path string, handler HandlerFunc) {
	s.Route("GET", path, handler, s.filters...)
}

func (s *Noc) POST(path string, handler HandlerFunc) {
	s.Route("POST", path, handler, s.filters...)
}

var logo string = ` _______                 
 \      \   ____   ____  
 /   |   \ /  _ \_/ ___\ 
/    |    (  <_> )  \___ 
\____|__  /\____/ \___  >
        \/            \/ `

func (s *Noc) Start(address string) error {
	fmt.Println(logo)
	logger.Logger.Infof("Noc Server %s is listening on %s", s.ServerName, address)
	return http.ListenAndServe(address, s.router)
}

func (s *Noc) AddFilters(filters ...FilterFunc) {
	s.filters = append(s.filters, filters...)
}

func NewNoc(serverName string) *Noc {
	routerDefault := NewMapRouter()

	return &Noc{
		ServerName: serverName,
		router:     routerDefault,
		filters:    make([]FilterFunc, 0, 4),
	}
}
