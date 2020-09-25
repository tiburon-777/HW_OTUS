package event

import (
	"time"
)

type Event struct {
	Title      string
	Date       time.Time
	Latency    time.Duration
	Note       string
	UserID     int64
	NotifyTime time.Duration
}
