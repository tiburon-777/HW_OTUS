package public

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

func (s Service) Create(ctx context.Context, e *CreateReq) (*CreateRsp, error) {
	var res CreateRsp
	ce, err := s.buildStorageEvent(e)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	t, err := s.App.Storage.Create(ce)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "storage error: can't create event")
	}
	res.ID = int64(t)
	return &res, nil
}

func (s Service) Update(ctx context.Context, e *UpdateReq) (*empty.Empty, error) {
	cid, ce, err := s.buildStorageEventAndID(e)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	if s.App.Storage.Update(cid, ce) != nil {
		return nil, status.Errorf(codes.Internal, "storage error: can't update event")
	}
	return nil, nil
}

func (s Service) Delete(ctx context.Context, e *DeleteReq) (*empty.Empty, error) {
	if s.App.Storage.Delete(event.ID(e.ID)) != nil {
		return nil, status.Errorf(codes.Internal, "storage error: can't update event")
	}
	return nil, nil
}

func (s Service) List(ctx context.Context, e *empty.Empty) (*ListResp, error) {
	tmp, err := s.App.Storage.List()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "storage error: can't get list of events")
	}
	l, err := s.buildEventList(tmp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	return &ListResp{Events: l}, nil
}

func (s Service) GetByID(ctx context.Context, e *GetByIDReq) (*GetByIDResp, error) {
	tmp, ok := s.App.Storage.GetByID(event.ID(e.ID))
	if !ok {
		return nil, status.Errorf(codes.NotFound, "event not found")
	}
	l, err := s.buildEventList(map[event.ID]event.Event{event.ID(e.ID): tmp})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	return &GetByIDResp{Events: l}, nil
}

func (s Service) GetByDate(ctx context.Context, e *GetByDateReq) (*GetByDateResp, error) {
	d, r, err := s.buildTimeAndRange(e)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	tmp, err := s.App.Storage.GetByDate(d, r)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "storage error: can't get list of events")
	}
	l, err := s.buildEventList(tmp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "inconvertible")
	}
	return &GetByDateResp{Events: l}, nil
}
