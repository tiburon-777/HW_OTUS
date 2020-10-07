package main

import (
	"flag"
	oslog "log"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/grpc"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/server/http"
	store "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage"
	googrpc "google.golang.org/grpc"
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

	listnGrpc, err := net.Listen("tcp", net.JoinHostPort(conf.Grpc.Address, conf.Grpc.Port))
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	serverGRPC := googrpc.NewServer()
	grpc.RegisterGrpcServer(serverGRPC, grpc.Service{App: *calendar})
	if err := serverGRPC.Serve(listnGrpc); err != nil {
		log.Errorf("failed to start grpc server: " + err.Error())
		os.Exit(1)
	}

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals)

		<-signals
		signal.Stop(signals)

		if err := serverHTTP.Stop(); err != nil {
			log.Errorf("failed to stop http server: " + err.Error())
		}
	}()

	if err := serverHTTP.Start(); err != nil {
		log.Errorf("failed to start http server: " + err.Error())
		os.Exit(1)
	}
}
