package log

import (
	"fmt"
	"time"

	"github.com/johnweldon/go/util"
)

type TimeRecord struct {
	Id             util.UUID
	Begin          time.Time
	Duration       time.Duration
	DurationString string
	Project        string
	Notes          string
	Tags           []string
}

func NewTimeRecord() TimeRecord {
	return TimeRecord{util.NewUUID(), time.Now(), 0, "", "", "", []string{}}
}

func (r *TimeRecord) SetEndPartial(t time.Time, fraction float64) error {
	if t.Before(r.Begin) {
		return fmt.Errorf("end %v is before begin %v", t, r.Begin)
	}
	r.Duration = time.Duration(float64(t.Sub(r.Begin)) * fraction)
	r.DurationString = r.Duration.String()
	return nil
}

func (r *TimeRecord) SetEnd(t time.Time) error {
	return r.SetEndPartial(t, 1.0)
}

