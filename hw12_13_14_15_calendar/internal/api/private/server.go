package private

import (
	"net"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/calendar"
	googrpc "google.golang.org/grpc"
)

type Server struct {
	s   *googrpc.Server
	app calendar.App
}

type Config struct {
	Address string
	Port    string
}

func New(app *calendar.App) Server {
	return Server{s: googrpc.NewServer(), app: *app}
}

func (s *Server) Start(conf Config) error {
	s.app.Logger.Infof("private GRPC server starting")
	listnGrpc, err := net.Listen("tcp", net.JoinHostPort(conf.Address, conf.Port))
	RegisterGrpcServer(s.s, &Service{})
	if err != nil {
		return err
	}
	return s.s.Serve(listnGrpc)
}

func (s *Server) Stop() {
	s.s.GracefulStop()
}
