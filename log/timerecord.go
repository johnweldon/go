package log

import (
	"time"
)

type TimeRecord struct {
	Begin    time.Time
	Duration time.Duration
	Project  string
	Notes    string
	Tags     []string
}
