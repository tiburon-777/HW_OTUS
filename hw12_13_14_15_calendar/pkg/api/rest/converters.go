package rest

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
)

func createReq2Event(e public.CreateReq) (res event.Event, err error) {
	res = event.Event{
		Title:  e.Title,
		Note:   e.Note,
		UserID: e.UserID,
	}
	res.Date, err = ptypes.Timestamp(e.Date)
	if err != nil {
		return event.Event{}, fmt.Errorf("can;t convert types: %w", err)
	}
	res.Latency, err = ptypes.Duration(e.Latency)
	if err != nil {
		return event.Event{}, fmt.Errorf("can;t convert types: %w", err)
	}
	res.NotifyTime, err = ptypes.Duration(e.NotifyTime)
	if err != nil {
		return event.Event{}, fmt.Errorf("can;t convert types: %w", err)
	}
	return res, nil
}

func pubEvent2Event(e public.Event) (res event.Event, err error) {
	res = event.Event{
		Title:  e.Title,
		Note:   e.Note,
		UserID: e.UserID,
	}
	res.Date, err = ptypes.Timestamp(e.Date)
	if err != nil {
		return event.Event{}, fmt.Errorf("can;t convert types: %w", err)
	}
	res.Latency, err = ptypes.Duration(e.Latency)
	if err != nil {
		return event.Event{}, fmt.Errorf("can;t convert types: %w", err)
	}
	res.NotifyTime, err = ptypes.Duration(e.NotifyTime)
	if err != nil {
		return event.Event{}, fmt.Errorf("can;t convert types: %w", err)
	}
	return res, nil
}

func event2pubEvent(e event.Event) (res public.Event, err error) {
	res = public.Event{
		Title:      e.Title,
		Latency:    ptypes.DurationProto(e.Latency),
		Note:       e.Note,
		UserID:     e.UserID,
		NotifyTime: ptypes.DurationProto(e.NotifyTime),
	}
	res.Date, err = ptypes.TimestampProto(e.Date)
	if err != nil {
		return public.Event{}, fmt.Errorf("can;t convert types: %w", err)
	}
	return res, nil
}

func events2pubEvents(e map[event.ID]event.Event) (res []public.Event, err error) {
	for id, ev := range e {
		r := public.Event{
			ID:         int64(id),
			Title:      ev.Title,
			Latency:    ptypes.DurationProto(ev.Latency),
			Note:       ev.Note,
			UserID:     ev.UserID,
			NotifyTime: ptypes.DurationProto(ev.NotifyTime),
		}
		r.Date, err = ptypes.TimestampProto(ev.Date)
		if err != nil {
			return []public.Event{}, fmt.Errorf("can;t convert types: %w", err)
		}
		res = append(res, r)
	}
	return res, nil
}
