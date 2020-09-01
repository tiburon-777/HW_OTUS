package store

import (
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/event"
	memorystorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/memory"
	sqlstorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/sql"
)

type Interface interface {
	Save(event event.Event) error
	Update(event event.Event) error
	Delete(event event.Event) error
	List() []event.Event
	Get(id string) (event.Event, bool)
}

func NewStore(config config.Config) Interface {
	if config.Storage.InMemory {
		st := memorystorage.New()
		return st
	}
	return sqlstorage.New()
}
