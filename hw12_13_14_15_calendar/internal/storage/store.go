package store

import (
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
	memorystorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/sql"
)

type Config struct {
	InMemory bool
	SQLHost  string
	SQLPort  string
	SQLDbase string
	SQLUser  string
	SQLPass  string
}

type StorageInterface interface {
	Create(event event.Event) (event.ID, error)
	Update(id event.ID, event event.Event) error
	Delete(id event.ID) error
	List() (map[event.ID]event.Event, error)
	GetByID(id event.ID) (event.Event, bool)
}

func NewStore(conf Config) StorageInterface {
	if conf.InMemory {
		st := memorystorage.New()
		return st
	}
	sqlConf := sqlstorage.Config{
		Host:  conf.SQLHost,
		Port:  conf.SQLPort,
		Dbase: conf.SQLDbase,
		User:  conf.SQLUser,
		Pass:  conf.SQLPass,
	}
	return sqlstorage.New(sqlConf)
}
