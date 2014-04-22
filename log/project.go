package log

import (
	"fmt"
)

type LegacyProject struct {
	TimeRecords  map[string]LegacyTimeRecord
	Name         string
	ProjectTotal float32
	ClockTime    float32
	Report       string
}

func (p LegacyProject) String() string {
	return fmt.Sprintf("%d records for a total %6.2f hours\n", len(p.TimeRecords), p.ProjectTotal)
}
