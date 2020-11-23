package public

import (
	"net"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/calendar"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
	googrpc "google.golang.org/grpc"
)

type Server struct {
	s   *googrpc.Server
	app calendar.App
}

func New(app *calendar.App) Server {
	return Server{s: googrpc.NewServer(), app: *app}
}

func (s *Server) Start(conf config.Calendar) error {
	s.app.Logger.Infof("GRPC server starting")
	listnGrpc, err := net.Listen("tcp", net.JoinHostPort(conf.GRPC.Address, conf.GRPC.Port))
	RegisterGrpcServer(s.s, &Service{})
	if err != nil {
		return err
	}
	return s.s.Serve(listnGrpc)
}

func (s *Server) Stop() {
	s.s.GracefulStop()
}
