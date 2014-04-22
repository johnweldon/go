package log

import ()

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
