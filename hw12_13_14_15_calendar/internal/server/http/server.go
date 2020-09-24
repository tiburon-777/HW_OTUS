package http

import (
	"context"
	"net"
	"net/http"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
)

type Server struct {
	server *http.Server
	app    app.App
	ctx    context.Context
	cancel context.CancelFunc
}

func NewServer(app *app.App, address string, port string) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{server: &http.Server{Addr: net.JoinHostPort(address, port), Handler: app.LoggingMiddleware(app.Handler)}, app: *app, ctx: ctx, cancel: cancel}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}
	s.app.Logger.Infof("Server starting")
	return nil
}

func (s *Server) Stop() error {
	if err := s.server.Close(); err != nil {
		return err
	}
	s.app.Logger.Infof("Server stoped")
	return nil
}
