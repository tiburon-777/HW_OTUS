package store

import (
	"time"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
	memorystorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/memory"
	sqlstorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/sql"
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
	Create(event.Event) (event.ID, error)
	Update(event.ID, event.Event) error
	Delete(event.ID) error
	List() (map[event.ID]event.Event, error)
	GetByID(event.ID) (event.Event, bool)
	GetByDate(time.Time, string) (map[event.ID]event.Event, error)
	GetNotifications() (map[event.ID]event.Event, error)
	SetNotified(event.ID) error
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
