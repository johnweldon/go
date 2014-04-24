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
			records = append(records, convertTimeRecord(v))
			if _, ok := working[v.ProjectName]; !ok {
				working[v.ProjectName] = NewProject(v.ProjectName)
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
	duration := end.Sub(begin)
	duration = time.Duration(float64(duration) * float64(r.Fraction))
	tr.Begin = begin
	tr.Duration = duration
	tr.Project = r.ProjectName
	tr.Notes = r.Notes
	tr.Tags = strings.Split(r.ProjectName, " ")
	return tr
}
