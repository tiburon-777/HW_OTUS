package sqlstorage

import (
	"context"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/event"
)

type Config struct {
	Host  string
	Port  string
	Dbase string
	User  string
	Pass  string
}

type Storage struct {
	//Events []event.Event
	//Mu     sync.RWMutex
}

func New(conf Config) *Storage {
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

func (s *Storage) Create(e event.Event) error {
	// TODO
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
