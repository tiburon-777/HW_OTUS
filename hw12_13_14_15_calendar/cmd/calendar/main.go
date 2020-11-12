package main

import (
	"context"
	"flag"
	"github.com/dmitryt/otus-golang-hw/hw12_13_14_15_calendar/service/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	oslog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/grpcserver"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/server/http"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
	flag.Parse()
}

func main() {
	conf, err := config.NewConfig(configFile)
	if err != nil {
		oslog.Fatal("не удалось открыть файл конфигурации:", err.Error())
	}
	log, err := logger.New(conf)
	if err != nil {
		oslog.Fatal("не удалось запустить логер:", err.Error())
	}
	storeConf := store.Config{
		InMemory: conf.Storage.InMemory,
		SQLHost:  conf.Storage.SQLHost,
		SQLPort:  conf.Storage.SQLPort,
		SQLDbase: conf.Storage.SQLDbase,
		SQLUser:  conf.Storage.SQLUser,
		SQLPass:  conf.Storage.SQLPass,
	}
	st := store.NewStore(storeConf)

	calendar := app.New(log, st)

	serverHTTP := internalhttp.NewServer(calendar, conf.Server.Address, conf.Server.Port)
	go func() {
		if err := serverHTTP.Start(); err != nil {
			log.Errorf("failed to start http server: " + err.Error())
			os.Exit(1)
		}
	}()

	serverGRPC := grpcserver.New(calendar)
	go func() {
		if err := serverGRPC.Start(conf); err != nil {
			log.Errorf("failed to start grpc server: " + err.Error())
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

	if err = server.RegisterCalendarHandler(context.Background(), grpcGwRouter, grpcDiler); err != nil {
		log.Errorf("can't register handlers for grpc-gateway: " + err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcGwRouter)
	go func() {
		log.Infof("start webAPI server")
		if err := http.ListenAndServe(net.JoinHostPort(conf.HTTP.Address, conf.HTTP.Port), mux); err != nil {
			log.Errorf("failed to start webAPI server: " + err.Error())
			os.Exit(1)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals,syscall.SIGINT, syscall.SIGHUP)
	<-signals
	signal.Stop(signals)
	serverGRPC.Stop()
	if err := serverHTTP.Stop(); err != nil {
		log.Errorf("failed to stop http server: " + err.Error())
	}
}
