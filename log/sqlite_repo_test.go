package log_test

import (
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/johnweldon/go/log"
)

const dbpath string = "./test.db"

func TestInit(t *testing.T) {
	os.Remove(dbpath)

	db := log.NewRelDB(dbpath)
	defer os.Remove(dbpath)

	records := []log.TimeRecord{}

	r1 := log.NewTimeRecord()
	r1.Duration = time.Duration(15) * time.Minute
	r1.DurationString = r1.Duration.String()
	r1.Project = "Project One"
	r1.Notes = "All my Notes"
	r1.Tags = append(r1.Tags, "Hello")
	r1.Tags = append(r1.Tags, "World")
	records = append(records, r1)

	r2 := log.NewTimeRecord()
	r2.Duration = time.Duration(3) * time.Hour
	r2.DurationString = r2.Duration.String()
	r2.Project = "Project Two"
	r2.Notes = "Other Notes"
	r2.Tags = append(r2.Tags, "Goodbye")
	r2.Tags = append(r2.Tags, "World")
	records = append(records, r2)

	err := db.SaveRecords(records)
	if err != nil {
		t.Error(err)
	}

	res := db.GetRecords()
	if len(res) != len(records) {
		t.Fail()
	}

	for _, rec := range res {
		t.Log(rec)
	}
}
