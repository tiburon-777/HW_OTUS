package grpc

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

type Service struct {
	App app.App
}

func (s Service) Create(ctx context.Context, e *Event) (*EventID, error) {
	var res EventID
	ce, err := pbevent2event(e)
	if err != nil {
		return nil, err
	}
	t, err := s.App.Storage.Create(ce)
	if err != nil {
		return nil, err
	}
	res.ID = int64(t)
	return &res, nil
}

func (s Service) Update(ctx context.Context, e *EventWthID) (*empty.Empty, error) {
	cid, ce, err := pbeventWitID2eventAndID(e)
	if err != nil {
		return nil, err
	}
	return nil, s.App.Storage.Update(cid, ce)
}

// TODO: Implement
func (s Service) Delete(ctx context.Context, e *EventID) (*empty.Empty, error) {
	return nil, s.App.Storage.Delete(event.ID(e.ID))
}

// TODO: Implement
func (s Service) List(ctx context.Context, e *empty.Empty) (*EventList, error) {
	tmp, err := s.App.Storage.List()
	if err != nil {
		return nil, err
	}
	return evtMap2pbEventList(tmp)
}

// TODO: Implement
func (s Service) GetByID(ctx context.Context, e *EventID) (*EventList, error) {
	tmp, ok := s.App.Storage.GetByID(event.ID(e.ID))
	if !ok {
		return nil, errors.New("event not found")
	}
	return evtMap2pbEventList(map[event.ID]event.Event{event.ID(e.ID): tmp})
}

// TODO: Implement
func (s Service) GetByDate(ctx context.Context, e *Date) (*EventList, error) {
	d, r, err := pbDate2Time(e)
	if err != nil {
		return nil, err
	}
	tmp, err := s.App.Storage.GetByDate(d, r)
	if err != nil {
		return nil, err
	}
	return evtMap2pbEventList(tmp)
}
