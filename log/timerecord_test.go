package log_test

import (
	"testing"
	"time"

	"github.com/johnweldon/go/log"
)

func Test(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2014-10-10")
	if err != nil {
		t.Errorf("error parsing date")
	}

	begin, err := time.Parse("15:04", "11:30")
	if err != nil {
		t.Errorf("error parsing begin time")
	}

	end, err := time.Parse("15:04", "12:45")
	if err != nil {
		t.Errorf("error parsing end time")
	}

	rec := log.NewTimeRecord()
	rec.Begin = begin
	rec.SetEnd(end)
	rec.SetDate(date)
	rec.Project = "TestProj"
	rec.Notes = "Test Notes"
	rec.Tags = append(rec.Tags, "Tag1")
	rec.Tags = append(rec.Tags, "Tag2")

	if rec.Begin.Year() != 2014 {
		t.Errorf("%s\n", rec)
	}
	if rec.Begin.Month() != time.October {
		t.Errorf("%s\n", rec)
	}
	if rec.Begin.Day() != 10 {
		t.Errorf("%s\n", rec)
	}

	if rec.DurationString != "1h15m0s" {
		t.Errorf("%s\n", rec)
	}

	if rec.String() != "2014-10-10 11:30 1h15m0s (TestProj) Test Notes ##[Tag1 Tag2]##" {
		t.Errorf(rec.String())
	}

}
