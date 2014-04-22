package log

import ()

type Project struct {
	TimeRecords  map[string]TimeRecord
	Name         string
	ProjectTotal float32
	ClockTime    float32
	Report       string
}
