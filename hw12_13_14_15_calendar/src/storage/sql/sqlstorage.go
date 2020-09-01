package sqlstorage

import (
	"context"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/event"
	"sync"
)

type Storage struct {
	Events []event.Event
	Mu     sync.RWMutex
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context) error {
	// TODO
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	// TODO
	return nil
}

func (s Storage) Save(e event.Event) error {
	if _, ok := s.Get(e.ID); !ok {
		s.Mu.Lock()
		s.Events = append(s.Events, e)
		s.Mu.Unlock()
	}
	return nil
}

func (s Storage) Update(event event.Event) error {
	return nil
}

func (s Storage) Delete(event event.Event) error {
	return nil
}

func (s Storage) List() []event.Event {
	return []event.Event{}
}

func (s Storage) Get(id string) (event.Event, bool) {
	return event.Event{}, false
}
