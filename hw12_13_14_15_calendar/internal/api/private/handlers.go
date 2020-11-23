package private

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/calendar"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	App calendar.App
}

func (s Service) GetNotifications(ctx context.Context, e *empty.Empty) (*GetRsp, error) {
	tmp, err := s.App.Storage.GetNotifications()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "storage error: can't get list of events")
	}
	l, err := s.buildEventList(tmp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	return &GetRsp{Events: l}, nil
}

func (s Service) SetNotified(ctx context.Context, r *SetReq) (*empty.Empty, error) {
	return nil, s.App.Storage.SetNotified(event.ID(r.ID))
}
