package log

import (
	"encoding/json"
	"io"
)

func ImportReportFromJson(r io.Reader) (*LegacyReport, error) {
	var report LegacyReport
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&report); err != nil {
		return nil, err
	}
	return &report, nil
}
