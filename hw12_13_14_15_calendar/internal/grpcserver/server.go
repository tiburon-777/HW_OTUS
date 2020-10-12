package grpcserver

import (
	"net"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/config"
	googrpc "google.golang.org/grpc"
)

type Server struct {
	s   *googrpc.Server
	app app.App
}

func New(app *app.App) Server {
	return Server{s: googrpc.NewServer(), app: *app}
}

func (s *Server) Start(conf config.Config) error {
	s.app.Logger.Infof("GRPC server starting")
	listnGrpc, err := net.Listen("tcp", net.JoinHostPort(conf.Grpc.Address, conf.Grpc.Port))
	RegisterGrpcServer(s.s, &Service{})
	if err != nil {
		return err
	}
	return s.s.Serve(listnGrpc)
}

func (s *Server) Stop() {
	s.s.GracefulStop()
}
