package log

import (
	"fmt"
)

type Report struct {
	ReportDate     string
	Projects       map[string]Project
	Notes          []string
	UnmatchedLines []string
	TimeRecords    []TimeRecord
}

func (r Report) String() string {
	return fmt.Sprintf("Report (%d Projects)\n%s\n", len(r.Projects), r.Projects)
}
