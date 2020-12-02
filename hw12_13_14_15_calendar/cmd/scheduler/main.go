package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/sheduler"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "", "Path to configuration file")
	flag.Parse()
}

func main() {
	var conf sheduler.Config
	err := config.New(configFile, &conf)
	if err != nil {
		log.Fatal("can't get config:", err.Error())
	}
	app := sheduler.New(conf)

	if err = app.Start(); err != nil {
		app.Logger.Errorf("failed to start scheduler: ", err.Error())
		os.Exit(1)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	<-signals
	signal.Stop(signals)
	app.Stop()
	log.Println("scheduler shutdown gracefully")
}
