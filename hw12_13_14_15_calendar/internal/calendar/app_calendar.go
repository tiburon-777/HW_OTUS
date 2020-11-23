package calendar

import (
	"context"
	"net/http"
	"time"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/logger"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage"
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

func (a *App) Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello! I'm calendar app!"))
}

func (a *App) LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			var path, useragent string
			if r.URL != nil {
				path = r.URL.Path
			}
			if len(r.UserAgent()) > 0 {
				useragent = r.UserAgent()
			}
			latency := time.Since(start)
			a.Logger.Infof("receive %s request from IP: %s on path: %s, duration: %s useragent: %s ", r.Method, r.RemoteAddr, path, latency, useragent)
		}()
		next.ServeHTTP(w, r)
	})
}
