package private

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/storage/event"
)

func (s Service) buildEventList(evtMap map[event.ID]event.Event) ([]*Event, error) {
	events := make([]*Event, len(evtMap))
	var err error
	var i int
	for k, v := range evtMap {
		evt := Event{ID: int64(k), Title: v.Title, UserID: v.UserID}
		evt.Date, err = ptypes.TimestampProto(v.Date)
		if err != nil {
			return nil, fmt.Errorf("can't convert date: %w", err)
		}
		events[i] = &evt
		i++
	}
	return events, nil
}
