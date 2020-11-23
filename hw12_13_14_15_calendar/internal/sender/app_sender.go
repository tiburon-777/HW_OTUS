package sender

import (
	"context"
	"encoding/json"
	oslog "log"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/api/private"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/logger"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/rabbit"
)

type Sender struct {
	Logger logger.Interface
	Rabbit *rabbit.Rabbit
	Queue  string
}

type Config struct {
	Rabbitmq config.Rabbit
	Logger   config.Logger
}

func New(conf Config) Sender {
	log, err := logger.New(logger.Config(conf.Logger))
	if err != nil {
		oslog.Fatal("не удалось запустить логер:", err.Error())
	}
	rb, err := rabbit.Attach(rabbit.Config(conf.Rabbitmq))
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ:", err.Error())
	}
	return Sender{Logger: log, Rabbit: rb, Queue: conf.Rabbitmq.Queue}
}

func (s *Sender) Start(ctx context.Context) error {
	msg, err := s.Rabbit.Consume(ctx, s.Queue)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case m := <-msg:
				var data []private.Event
				err := json.Unmarshal(m.Data, &data)
				if err != nil {
					s.Logger.Errorf("can`t unmarshal data", err)
				}
				for _, v := range data {
					s.Logger.Infof("User %s notified about event %s", v.UserID, v.ID)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (s *Sender) Stop(cancel context.CancelFunc) {
	cancel()
}
