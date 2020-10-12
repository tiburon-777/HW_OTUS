package grpcserver

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

func (s Service) buildStorageEvent (pbe *CreateReq) (res event.Event, err error) {
	res = event.Event{Title: pbe.Title, Note: pbe.Note, UserID: pbe.UserID}
	res.Date, err = ptypes.Timestamp(pbe.Date)
	if err != nil {
		return event.Event{}, err
	}
	res.Latency, err = ptypes.Duration(pbe.Latency)
	if err != nil {
		return event.Event{}, err
	}
	res.NotifyTime, err = ptypes.Duration(pbe.NotifyTime)
	if err != nil {
		return event.Event{}, err
	}
	return res, nil
}

func (s Service) buildStorageEventAndID(pbe *UpdateReq) (id event.ID, evt event.Event, err error) {
	evt = event.Event{Title: pbe.Event.Title, Note: pbe.Event.Note, UserID: pbe.Event.UserID}
	evt.Date, err = ptypes.Timestamp(pbe.Event.Date)
	if err != nil {
		return 0, event.Event{}, err
	}
	evt.Latency, err = ptypes.Duration(pbe.Event.Latency)
	if err != nil {
		return 0, event.Event{}, err
	}
	evt.NotifyTime, err = ptypes.Duration(pbe.Event.NotifyTime)
	if err != nil {
		return 0, event.Event{}, err
	}
	return event.ID(pbe.ID), evt, nil
}


func (s Service) buildEventList(evtMap map[event.ID]event.Event) ([]*Event, error) {
	var events = []*Event{}
	var err error
	for k, v := range evtMap {
		evt := Event{ID: int64(k), Title: v.Title, Latency: ptypes.DurationProto(v.Latency), Note: v.Note, UserID: v.UserID, NotifyTime: ptypes.DurationProto(v.NotifyTime)}
		evt.Date, err = ptypes.TimestampProto(v.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, &evt)
	}
	return events, err
}

func (s Service) buildTimeAndRange(e *GetByDateReq) (start time.Time, qrange string, err error) {
	date, err := ptypes.Timestamp(e.Date)
	return date, string(e.Range), err
}
