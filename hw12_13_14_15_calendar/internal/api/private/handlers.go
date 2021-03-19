package private

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/calendar"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
	googrpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	S   *googrpc.Server
	App calendar.App
}

type Config struct {
	Address string
	Port    string
}

func New(app *calendar.App) Service {
	return Service{S: googrpc.NewServer(), App: *app}
}

func (s *Service) Start(conf Config) error {
	s.App.Logger.Infof("private GRPC server starting")
	listnGrpc, err := net.Listen("tcp", net.JoinHostPort(conf.Address, conf.Port))
	RegisterGrpcServer(s.S, s)
	if err != nil {
		return fmt.Errorf("can't start private GRPC service: %w", err)
	}
	return s.S.Serve(listnGrpc)
}

func (s *Service) Stop() {
	s.S.GracefulStop()
}

func (s Service) GetNotifications(ctx context.Context, e *empty.Empty) (*GetRsp, error) {
	tmp, err := s.App.Storage.GetNotifications()
	if err != nil {
		s.App.Logger.Errorf("storage error: can't get list of events: %w", err)
		return nil, status.Errorf(codes.Internal, "storage error: can't get list of events")
	}
	l, err := s.buildEventList(tmp)
	if err != nil {
		s.App.Logger.Errorf("inconvertible: %w", err)
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	return &GetRsp{Events: l}, nil
}

func (s Service) SetNotified(ctx context.Context, r *SetReq) (*empty.Empty, error) {
	return nil, s.App.Storage.SetNotified(event.ID(r.ID))
}

func (s Service) PurgeOldEvents(ctx context.Context, r *PurgeReq) (*PurgeResp, error) {
	q, err := s.App.Storage.PurgeOldEvents(r.OlderThenDays)
	if err != nil {
		return &PurgeResp{}, fmt.Errorf("can't purge old events from storage: %w", err)
	}
	return &PurgeResp{Qty: q}, nil
}
