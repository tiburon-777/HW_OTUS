package main

import (
	"flag"
	oslog "log"
	"os"
	"os/signal"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/app"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/memory"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
	flag.Parse()
}

func main() {
	conf, err := config.NewConfig(configFile)
	if err != nil { oslog.Fatal("не удалось открыть файл конфигурации:", err.Error())}

	log, err := logger.New(conf)
	if err != nil { oslog.Fatal("не удалось запустить логер:", err.Error())}

	storage := memorystorage.New()
	calendar := app.New(log, storage)

	server := internalhttp.NewServer(calendar)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals)

		<-signals
		signal.Stop(signals)

		if err := server.Stop(); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
	}()

	if err := server.Start(); err != nil {
		log.Error("failed to start http server: " + err.Error())
		os.Exit(1)
	}
}
