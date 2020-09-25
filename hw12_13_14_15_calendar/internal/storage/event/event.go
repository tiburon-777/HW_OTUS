package event

import (
	"time"
)

type ID int64

type Event struct {
	Title      string
	Date       time.Time
	Latency    time.Duration
	Note       string
	UserID     int64
	NotifyTime time.Duration
}
