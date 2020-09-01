package memorystorage

import (
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/event"
)

type Storage struct {
	//events []event.Event
	//mu     sync.RWMutex
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Save(event event.Event) error {
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
