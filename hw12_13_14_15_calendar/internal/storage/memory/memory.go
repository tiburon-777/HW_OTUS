package memory

import (
	"sync"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

type Storage struct {
	Events map[event.ID]event.Event
	lastID event.ID
	Mu     sync.RWMutex
}

func New() *Storage {
	return &Storage{Events: make(map[event.ID]event.Event)}
}

func (s *Storage) Create(event event.Event) (event.ID, error) {
	s.Mu.Lock()
	s.lastID++
	s.Events[s.lastID] = event
	s.Mu.Unlock()
	return s.lastID, nil
}

func (s *Storage) Update(id event.ID, event event.Event) error {
	s.Mu.Lock()
	s.Events[id] = event
	s.Mu.Unlock()
	return nil
}

func (s *Storage) Delete(id event.ID) error {
	s.Mu.Lock()
	delete(s.Events, id)
	s.Mu.Unlock()
	return nil
}

func (s *Storage) List() (map[event.ID]event.Event, error) {
	return s.Events, nil
}

func (s *Storage) GetByID(id event.ID) (event.Event, bool) {
	if s.Events[id].Title == "" {
		return event.Event{}, false
	}
	return s.Events[id], true
}
