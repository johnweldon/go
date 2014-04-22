package log

import (
	"fmt"
)

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
