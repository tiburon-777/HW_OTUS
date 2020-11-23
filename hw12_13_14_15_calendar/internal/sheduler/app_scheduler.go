package sheduler

import (
	"context"
	oslog "log"
	"time"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/logger"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/rabbit"
)

type Scheduler struct {
	CalendarAPI config.Server
	Logger      logger.Interface
	Rabbit      *rabbit.Rabbit
}

type Config struct {
	CalendarAPI config.Server
	Rabbitmq    config.Rabbit
	Storage     config.Storage
	Logger      config.Logger
}

func New(conf Config) Scheduler {
	log, err := logger.New(logger.Config(conf.Logger))
	if err != nil {
		oslog.Fatal("не удалось запустить логер:", err.Error())
	}
	rb, err := rabbit.New(rabbit.Config(conf.Rabbitmq))
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ:", err.Error())
	}
	return Scheduler{CalendarAPI: conf.CalendarAPI, Logger: log, Rabbit: rb}
}

func (s *Scheduler) Start(ctx context.Context) error {
	fetcher := riseOnTick(ctx, s.Logger, func() interface{} { return worker(ctx, s.CalendarAPI, s.Rabbit) }, 1*time.Minute)
	go func() {
		for {
			select {
			case err := <-fetcher:
				s.Logger.Infof("fetcher closed")
				if err != nil {
					s.Logger.Errorf("...with error:", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (s *Scheduler) Stop(cancel context.CancelFunc) {
	cancel()
}
