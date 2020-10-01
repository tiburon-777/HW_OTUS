package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/grpc"
)

type Service struct{}

func (s Service) Create(context.Context, *grpc.Event) (*grpc.EventID, error) {
	return nil, nil
}

func (s Service) Update(context.Context, *grpc.EventWthID) (*empty.Empty, error) {
	return nil, nil
}

func (s Service) Delete(context.Context, *grpc.EventID) (*empty.Empty, error) {
	return nil, nil
}

func (s Service) List(context.Context, *empty.Empty) (*grpc.EventList, error) {
	return nil, nil
}

func (s Service) GetByID(context.Context, *grpc.EventID) (*grpc.EventList, error) {
	return nil, nil
}

func (s Service) GetByDate(context.Context, *grpc.Date) (*grpc.EventList, error) {
	return nil, nil
}
