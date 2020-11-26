package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/api/private"
	"google.golang.org/grpc"
	oslog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
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
	oslog.Printf("Переменная APP_GRPC_PORT: %#v", os.Getenv("APP_GRPC_PORT"))
	oslog.Printf("Конфиг приложения: %#v", conf)
	log, err := logger.New(logger.Config(conf.Logger))
	if err != nil {
		oslog.Fatal("не удалось запустить логер:", err.Error())
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

	grpcDiler, err := grpc.Dial(net.JoinHostPort(conf.HTTP.Address, conf.HTTP.Port), grpc.WithInsecure())
	if err != nil {
		log.Errorf("can't dial grpc server: " + err.Error())
		os.Exit(1)
	}
	defer grpcDiler.Close()

	grpcGwRouter := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	if err = public.RegisterGrpcHandler(ctx, grpcGwRouter, grpcDiler); err != nil {
		log.Errorf("can't register handlers for grpc-gateway: " + err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcGwRouter)
	go func() {
		log.Infof("webAPI server starting")
		if err := http.ListenAndServe(net.JoinHostPort(conf.HTTP.Address, conf.HTTP.Port), mux); err != nil {
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
