package memorystorage

import (
	"sync"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/event"
)

type Storage struct {
	Events []event.Event
	Mu     sync.RWMutex
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Create(event event.Event) error {
	if _, ok := s.Get(event.ID); !ok {
		s.Mu.Lock()
		s.Events = append(s.Events, event)
		s.Mu.Unlock()
	}
	return nil
}

func (s *Storage) Update(event event.Event) error {
	return nil
}

func (s *Storage) Delete(event event.Event) error {
	return nil
}

func (s *Storage) List() []event.Event {
	return []event.Event{}
}

func (s *Storage) Get(id string) (event.Event, bool) {
	return event.Event{}, false
}
