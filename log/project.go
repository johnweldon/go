package log

import (
	"fmt"
)

type Project struct {
	TimeRecords  map[string]TimeRecord
	Name         string
	ProjectTotal float32
	ClockTime    float32
	Report       string
}

func (p Project) String() string {
	return fmt.Sprintf("%d records for a total %6.2f hours\n", len(p.TimeRecords), p.ProjectTotal)
}
