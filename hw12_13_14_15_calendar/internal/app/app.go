package app

import (
	"context"
	"net/http"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/logger"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

type Interface interface {
	CreateEvent(ctx context.Context, title string) (err error)
	Handler(w http.ResponseWriter, r *http.Request)
	LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc
}

type App struct {
	Storage store.StorageInterface
	Logger  logger.Interface
}

func New(logger logger.Interface, storage store.StorageInterface) *App {
	return &App{Logger: logger, Storage: storage}
}

func (a *App) CreateEvent(ctx context.Context, title string) (err error) {
	_, err = a.Storage.Create(event.Event{Title: title})
	if err != nil {
		a.Logger.Errorf("can not create event")
	}
	a.Logger.Infof("event created")
	return err
}

func (a *App) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello! I'm calendar app!"))
}

func (a *App) LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.Logger.Infof("receive %s request from IP %s", r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
