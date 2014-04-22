package log

import ()

type Report struct {
	ReportDate     string
	Projects       map[string]Project
	Notes          []string
	UnmatchedLines []string
	TimeRecords    []TimeRecord
}
