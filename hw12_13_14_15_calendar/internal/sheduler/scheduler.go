package sheduler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/api/private"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/config"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/logger"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/rabbit"
)

func riseOnTick(ctx context.Context, log logger.Interface, fn func() interface{}, interval time.Duration) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		t := time.NewTicker(interval)
		defer func() {
			t.Stop()
			close(valueStream)
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				select {
				case <-ctx.Done():
					return
				default:
					log.Infof("scheduler come to rabbit")
					valueStream <- fn()
				}
			}
		}
	}()
	return valueStream
}

func worker(ctx context.Context, calendarAPI config.Server, rb *rabbit.Rabbit) error {
	cli, err := private.NewClient(calendarAPI.Address, calendarAPI.Port)
	if err != nil {
		return err
	}
	resp, err := cli.GetNotifications(ctx, nil)
	if err != nil {
		return err
	}
	for event := range resp.Events {
		b, err := json.Marshal(event)
		if err != nil {
			return err
		}
		err = rb.Publish(string(b))
		if err != nil {
			return err
		}
	}
	return nil
}