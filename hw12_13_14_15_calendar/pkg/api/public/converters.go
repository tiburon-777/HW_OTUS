package public

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
)

func (s Service) buildStorageEvent(pbe *CreateReq) (res event.Event, err error) {
	res = event.Event{Title: pbe.Title, Note: pbe.Note, UserID: pbe.UserID}
	res.Date, err = ptypes.Timestamp(pbe.Date)
	if err != nil {
		return event.Event{}, fmt.Errorf("can't convert date: %w", err)
	}
	res.Latency, err = ptypes.Duration(pbe.Latency)
	if err != nil {
		return event.Event{}, fmt.Errorf("can't convert duration: %w", err)
	}
	res.NotifyTime, err = ptypes.Duration(pbe.NotifyTime)
	if err != nil {
		return event.Event{}, fmt.Errorf("can't convert duration: %w", err)
	}
	return res, nil
}

func (s Service) buildStorageEventAndID(pbe *UpdateReq) (id event.ID, evt event.Event, err error) {
	evt = event.Event{Title: pbe.Event.Title, Note: pbe.Event.Note, UserID: pbe.Event.UserID}
	evt.Date, err = ptypes.Timestamp(pbe.Event.Date)
	if err != nil {
		return 0, event.Event{}, fmt.Errorf("can't convert time: %w", err)
	}
	evt.Latency, err = ptypes.Duration(pbe.Event.Latency)
	if err != nil {
		return 0, event.Event{}, fmt.Errorf("can't convert duration: %w", err)
	}
	evt.NotifyTime, err = ptypes.Duration(pbe.Event.NotifyTime)
	if err != nil {
		return 0, event.Event{}, fmt.Errorf("can't convert duration: %w", err)
	}
	return event.ID(pbe.ID), evt, nil
}

func (s Service) buildEventList(evtMap map[event.ID]event.Event) ([]*Event, error) {
	events := make([]*Event, len(evtMap))
	var err error
	var i int
	for k, v := range evtMap {
		evt := Event{ID: int64(k), Title: v.Title, Latency: ptypes.DurationProto(v.Latency), Note: v.Note, UserID: v.UserID, NotifyTime: ptypes.DurationProto(v.NotifyTime)}
		evt.Date, err = ptypes.TimestampProto(v.Date)
		if err != nil {
			return nil, fmt.Errorf("can't convert date: %w", err)
		}
		events[i] = &evt
		i++
	}
	return events, err
}

func (s Service) buildTimeAndRange(e *GetByDateReq) (start time.Time, qrange string, err error) {
	date, err := ptypes.Timestamp(e.Date)
	if err != nil {
		return time.Time{}, "", fmt.Errorf("can't convert date: %w", err)
	}
	return date, e.Range.String(), nil
}
