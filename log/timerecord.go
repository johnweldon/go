package log

import (
	"time"

	"github.com/johnweldon/go/util"
)

type TimeRecord struct {
	id       util.UUID
	Begin    time.Time
	Duration time.Duration
	Project  string
	Notes    string
	Tags     []string
}

func NewTimeRecord() TimeRecord {
	return TimeRecord{util.NewUUID(), time.Now(), 0, "", "", []string{}}
}
