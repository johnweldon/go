package log

import (
	"encoding/json"
	"io"
	"strings"
	"time"
)

var LegacyTimeFormat string = "2006-01-02T15:04:05"

func ImportReportFromJson(r io.Reader) (*LegacyReport, error) {
	var report LegacyReport
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&report); err != nil {
		return nil, err
	}
	return &report, nil
}

func ConvertLegacyRecords(r *LegacyReport) (records []TimeRecord, projects []Project) {
	working := map[string]Project{}
	if r != nil {
		for _, v := range r.TimeRecords {
			converted := convertTimeRecord(v)
			records = append(records, converted)
			if _, ok := working[converted.Project]; !ok {
				working[converted.Project] = NewProject(converted.Project)
			}
		}
	}
	for _, v := range working {
		projects = append(projects, v)
	}
	return
}

func convertTimeRecord(r LegacyTimeRecord) TimeRecord {
	tr := NewTimeRecord()
	begin, err := time.ParseInLocation(LegacyTimeFormat, r.BeginTime, time.Local)
	if err != nil {
		begin = time.Now()
	}
	end, err := time.ParseInLocation(LegacyTimeFormat, r.EndTime, time.Local)
	if err != nil {
		end = time.Now()
	}
	tr.Begin = begin
	tr.SetEndPartial(end, float64(r.Fraction))
	tr.Notes = r.Notes
	tr.Tags = strings.Split(r.ProjectName, " ")
	if len(tr.Tags) > 0 {
		tr.Project = tr.Tags[0]
	}
	return tr
}
