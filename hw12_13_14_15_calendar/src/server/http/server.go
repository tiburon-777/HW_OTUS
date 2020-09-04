package internalhttp

import (
	"net"
	"net/http"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/app"
)

type Server struct {
	server *http.Server
	app    app.Interface
}

func NewServer(app app.Interface, address string, port string) *Server {
	return &Server{server: &http.Server{Addr: net.JoinHostPort(address, port), Handler: http.HandlerFunc(app.Handler)}, app: app}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Close()
}
