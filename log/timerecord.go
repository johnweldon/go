package log

import (
	"fmt"
	"time"

	"github.com/johnweldon/go-misc/util"
)

type TimeRecord struct {
	ID             util.UUID `bson:"_id"`
	Begin          time.Time
	Duration       time.Duration
	DurationString string
	Project        string
	Notes          string
	Tags           []string
}

var _ fmt.Stringer = (*TimeRecord)(nil)

func NewTimeRecord() TimeRecord {
	return TimeRecord{util.NewUUID(), time.Now(), 0, "", "", "", []string{}}
}

func (r *TimeRecord) SetDate(date time.Time) {
	r.Begin = time.Date(date.Year(), date.Month(), date.Day(), r.Begin.Hour(), r.Begin.Minute(), r.Begin.Second(), 0, date.Location())
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

func (r *TimeRecord) String() string {
	return fmt.Sprintf("%s %s (%s) %s ##%s##", r.Begin.Format("2006-01-02 15:04"), r.DurationString, r.Project, r.Notes, r.Tags)
}
