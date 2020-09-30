package memory

import (
	"sync"
	"time"

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

func (s *Storage) GetByDate(startDate time.Time, rng string) (map[event.ID]event.Event, error) {
	endDate := getEndDate(startDate, rng)
	s.Mu.Lock()
	defer s.Mu.Unlock()
	res := make(map[event.ID]event.Event)
	for k, v := range s.Events {
		if (v.Date.After(startDate) && v.Date.Before(endDate)) ||
			(v.Date.Add(v.Latency).Before(v.Date) && v.Date.Add(v.Latency).After(v.Date)) ||
			(v.Date.Before(startDate)) && (v.Date.Add(v.Latency).After(endDate)) {
			res[k] = v
		}
	}
	return res, nil
}

func getEndDate(startDate time.Time, rng string) time.Time {
	switch rng {
	case "DAY":
		return startDate.AddDate(0, 0, 1)
	case "WEEK":
		return startDate.AddDate(0, 0, 7)
	case "MONTH":
		return startDate.AddDate(0, 1, 0)
	default:
		return startDate
	}
}
