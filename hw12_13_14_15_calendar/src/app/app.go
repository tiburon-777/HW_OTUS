package app

import (
	"context"
	"net/http"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/logger"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/src/storage/event"
)

type Interface interface {
	CreateEvent(ctx context.Context, id string, title string) (err error)
	Handler(w http.ResponseWriter, r *http.Request)
}

type App struct {
	storage store.StorageInterface
	logger  logger.Interface
}

func New(logger logger.Interface, storage store.StorageInterface) *App {
	return &App{logger: logger, storage: storage}
}

func (a *App) CreateEvent(ctx context.Context, id string, title string) (err error) {
	err = a.storage.Create(event.Event{ID: id, Title: title})
	if err != nil {
		a.logger.Errorf("can not create event")
	}
	a.logger.Infof("event created")
	return err
}

func (a *App) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello! I'm calendar app!"))
}
