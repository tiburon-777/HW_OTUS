package grpc

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/internal/storage/event"
)

func pbevent2event(pbe *Event) (res event.Event, err error) {
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

func pbeventWitID2eventAndID(pbe *EventWthID) (id event.ID, evt event.Event, err error) {
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
