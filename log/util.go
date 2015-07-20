package log

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"
	"time"
)

const (
	LegacyTimeFormat string = "2006-01-02T15:04:05"

	CSVDateFormat   string = "1/2/2006"
	CSVTimeFormat   string = "03:04:05 PM"
	CSVDateIndex    int    = 0
	CSVBeginIndex   int    = 1
	CSVEndIndex     int    = 2
	CSVMessageIndex int    = 4
)

func ImportReportFromJSON(r io.Reader) (*LegacyReport, error) {
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

func ConvertCSVRecords(r io.Reader, project string) ([]TimeRecord, error) {
	res := []TimeRecord{}
	csvReader := csv.NewReader(r)
	records, err := csvReader.ReadAll()
	if err != nil {
		return res, err
	}
	for _, record := range records {
		if len(record) < 5 {
			continue
		}
		date, err := time.ParseInLocation(CSVDateFormat, record[CSVDateIndex], time.Local)
		if err != nil {
			continue
		}
		begin, err := time.ParseInLocation(CSVTimeFormat, record[CSVBeginIndex], time.Local)
		if err != nil {
			continue
		}
		end, err := time.ParseInLocation(CSVTimeFormat, record[CSVEndIndex], time.Local)
		if err != nil {
			continue
		}
		message := record[CSVMessageIndex]
		rec := NewTimeRecord()
		rec.Begin = begin
		rec.Notes = message
		rec.Project = project
		rec.SetEnd(end)
		rec.SetDate(date)

		res = append(res, rec)
	}
	return res, nil
}
