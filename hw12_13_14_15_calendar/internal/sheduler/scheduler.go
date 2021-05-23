package sheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
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

func worker(ctx context.Context, calendarAPI config.Server, rb *rabbit.Rabbit, log logger.Interface) error {
	cli, err := private.NewClient(calendarAPI.Address, calendarAPI.Port)
	if err != nil {
		return fmt.Errorf("can't get GRPC client: %w", err)
	}
	resp, err := cli.GetNotifications(ctx, &empty.Empty{})
	if err != nil {
		return fmt.Errorf("can't get events from GRPC endpoint: %w", err)
	}
	for _, event := range resp.Events {
		b, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("can't marshal events into JSON: %w", err)
		}
		err = rb.Publish(b)
		if err != nil {
			return fmt.Errorf("can't publish serialized data to RabbitMQ: %w", err)
		}
	}
	resp2, err := cli.PurgeOldEvents(ctx, &private.PurgeReq{OlderThenDays: 365})
	if err != nil {
		return fmt.Errorf("can't purge old events: %w", err)
	}
	log.Infof("Scheduler successfully purges %s events from storage", resp2.Qty)
	return nil
}
