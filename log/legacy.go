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

type LegacyReport struct {
	ReportDate     string
	Projects       map[string]LegacyProject
	Notes          []string
	UnmatchedLines []string
	TimeRecords    []LegacyTimeRecord
}

func (r LegacyReport) String() string {
	return fmt.Sprintf("LegacyReport (%d Projects)\n%s\n", len(r.Projects), r.Projects)
}

type LegacyTimeRecord struct {
	BeginTime            string
	EndTime              string
	Fraction             float32
	Notes                string
	ProjectName          string
	TimeSpanMinutes      float32
	TimeSpanQuarterHours float32
	TimeSpanHours        float32
	BillableHours        float32
	ReportLine           string
}
