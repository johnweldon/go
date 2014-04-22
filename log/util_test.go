package log

import (
	"strings"
	"testing"
)

var jsonStr string = `{
    "ReportDate": "2014-04-21",
    "Projects": {
        "PROJECT": {
            "TimeRecords": {
                "2014-04-21T08:00:00": {
                    "BeginTime": "2014-04-21T08:00:00",
                    "EndTime": "2014-04-21T09:00:00",
                    "Fraction": 1.0,
                    "Notes": "example notes line",
                    "ProjectName": "PROJECT",
                    "TimeSpanMinutes": 60.0,
                    "TimeSpanQuarterHours": 4.0,
                    "TimeSpanHours": 1.0,
                    "BillableHours": 1.0,
                    "ReportLine": "0800-0900 example notes line"
                }
            },
            "Name": "PROJECT",
            "ProjectTotal": 1.0,
            "ClockTime": 1.0,
            "Report": ""
        }
    },
    "Notes": [
        "example notes line"
    ],
    "UnmatchedLines": [
        "",
        "asdf",
        "qwer"
    ],
    "TimeRecords": [
        {
            "BeginTime": "2014-04-21T08:00:00",
            "EndTime": "2014-04-21T09:00:00",
            "Fraction": 1.0,
            "Notes": "example notes line",
            "ProjectName": "PROJECT",
            "TimeSpanMinutes": 60.0,
            "TimeSpanQuarterHours": 4.0,
            "TimeSpanHours": 1.0,
            "BillableHours": 1.0,
            "ReportLine": "0800-0900 example notes line"
        }
    ]
}`

func TestImportReportFromJson(t *testing.T) {
	report, err := ImportReportFromJson(strings.NewReader(jsonStr))
	if err != nil {
		t.Error(err)
	}
	t.Log(report)
}
