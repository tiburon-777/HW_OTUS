package store

import (
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/event"
	memorystorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/memory"
	sqlstorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/sql"
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
	Create(event event.Event) error
	Update(event event.Event) error
	Delete(event event.Event) error
	List() []event.Event
	Get(id string) (event.Event, bool)
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
