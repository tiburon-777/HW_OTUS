package main

import (
	"context"
	"flag"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/sender"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/sender.conf", "Path to configuration file")
	flag.Parse()
}

func main() {
	var conf sender.Config
	err := config.New(configFile, &conf)
	if err != nil {
		log.Fatal("не удалось открыть файл конфигурации:", err.Error())
	}
	app := sender.New(conf)

	ctx, cancel := context.WithCancel(context.Background())
	if err = app.Start(ctx); err != nil {
		app.Logger.Errorf("failed to start sender: ", err.Error())
		os.Exit(1)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	<-signals
	signal.Stop(signals)
	app.Stop(cancel)
}
