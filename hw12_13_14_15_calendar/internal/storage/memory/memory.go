package memory

import (
	"sync"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

type Storage struct {
	Events map[int64]event.Event
	lastID int64
	Mu     sync.RWMutex
}

func New() *Storage {
	return &Storage{Events: make(map[int64]event.Event)}
}

func (s *Storage) Create(event event.Event) (int64, error) {
	s.Mu.Lock()
	s.lastID++
	s.Events[s.lastID] = event
	s.Mu.Unlock()
	return s.lastID, nil
}

func (s *Storage) Update(id int64, event event.Event) error {
	s.Mu.Lock()
	s.Events[id] = event
	s.Mu.Unlock()
	return nil
}

func (s *Storage) Delete(id int64) error {
	s.Mu.Lock()
	delete(s.Events, id)
	s.Mu.Unlock()
	return nil
}

func (s *Storage) List() (map[int64]event.Event, error) {
	return s.Events, nil
}

func (s *Storage) GetByID(id int64) (event.Event, bool) {
	if s.Events[id].Title == "" {
		return event.Event{}, false
	}
	return s.Events[id], true
}
