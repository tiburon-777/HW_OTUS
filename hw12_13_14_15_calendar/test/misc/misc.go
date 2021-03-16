package misc

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"log"
	"time"
)

var TestEvent01 = public.CreateReq{
	Title:      "Test event 01",
	Date:       Time2pbtimestamp(time.Now().Add(30 * time.Second)),
	Latency:    Dur2pbduration(24 * time.Hour),
	Note:       "Note of test event 01",
	NotifyTime: Dur2pbduration(5 * time.Minute),
	UserID:     1111,
}

var TestEvent02 = public.CreateReq{
	Title:      "Test event 02",
	Date:       Time2pbtimestamp(time.Now().Add(60 * time.Second)),
	Latency:    Dur2pbduration(2 * 24 * time.Hour),
	Note:       "Note of test event 02",
	NotifyTime: Dur2pbduration(5 * time.Minute),
	UserID:     2222,
}

func Time2pbtimestamp(t time.Time) *timestamp.Timestamp {
	r, err := ptypes.TimestampProto(t)
	if err != nil {
		log.Fatalf("cant convert Time to Timestamp: %s", err.Error())
	}
	return r
}

func Dur2pbduration(t time.Duration) *duration.Duration {
	return ptypes.DurationProto(t)
}
