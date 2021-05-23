package main

import (
	"context"
	"flag"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/api/private"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/rest"
	oslog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/calendar"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/logger"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "", "Path to configuration file")
	flag.Parse()
}

func main() {
	var conf config.Calendar
	err := config.New(configFile, &conf)
	if err != nil {
		oslog.Fatal("can't get config:", err.Error())
	}
	log, err := logger.New(logger.Config(conf.Logger))
	if err != nil {
		oslog.Fatal("can't start logger:", err.Error())
	}

	st := store.NewStore(store.Config(conf.Storage))

	calendar := calendar.New(log, st)

	serverGRPC := public.New(calendar)
	go func() {
		if err := serverGRPC.Start(conf); err != nil {
			log.Errorf("failed to start grpc server: " + err.Error())
			os.Exit(1)
		}
	}()

	serverAPI := private.New(calendar)
	go func() {
		if err := serverAPI.Start(private.Config(conf.API)); err != nil {
			log.Errorf("failed to start API server: " + err.Error())
			os.Exit(1)
		}
	}()

	_, cancel := context.WithCancel(context.Background())
	m := mux.NewRouter()

	m.HandleFunc("/events", rest.FromRESTCreate(calendar)).Methods("POST")
	m.HandleFunc("/events/{ID}", rest.FromRESTUpdate(calendar)).Methods("PUT")
	m.HandleFunc("/events/{ID}", rest.FromRESTDelete(calendar)).Methods("DELETE")
	m.HandleFunc("/events", rest.FromRESTList(calendar)).Methods("GET")
	m.HandleFunc("/events/{ID}", rest.FromRESTGetByID(calendar)).Methods("GET")
	m.HandleFunc("/events/{Range}/{Date}", rest.FromRESTGetByDate(calendar)).Methods("GET")

	go func() {
		log.Infof("webAPI server starting")
		if err := http.ListenAndServe(net.JoinHostPort(conf.HTTP.Address, conf.HTTP.Port), m); err != nil {
			log.Errorf("failed to start webAPI server: " + err.Error())
			os.Exit(1)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	<-signals
	signal.Stop(signals)
	serverGRPC.Stop()
	serverAPI.Stop()
	cancel()
}
