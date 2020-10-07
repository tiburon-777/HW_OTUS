package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
)

type Service struct {
	App app.App
}

func (s Service) Create(ctx context.Context, e *Event) (*EventID, error) {
	var res EventID
	//var tmp = event.Event{e.Title, e.Date.(time.Time), e.Latency, e.Note, e.UserID, e.NotifyTime}
	//t, err := s.App.Storage.Create(tmp)
	//if err != nil { return nil, err }
	//res.ID = string(t)
	return &res, nil
}

func (s Service) Update(ctx context.Context, e *EventWthID) (*empty.Empty, error) {
	return nil, nil
}

func (s Service) Delete(ctx context.Context, e *EventID) (*empty.Empty, error) {
	return nil, nil
}

func (s Service) List(ctx context.Context, e *empty.Empty) (*EventList, error) {
	return nil, nil
}

func (s Service) GetByID(ctx context.Context, e *EventID) (*EventList, error) {
	return nil, nil
}

func (s Service) GetByDate(ctx context.Context, e *Date) (*EventList, error) {
	return nil, nil
}
