package log

import (
	"strings"
	"testing"
)

var jsonStr = `{
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
                },
                "2014-04-21T09:00:00": {
                    "BeginTime": "2014-04-21T09:00:00",
                    "EndTime": "2014-04-21T10:30:00",
                    "Fraction": 0.5,
                    "Notes": "example notes line 2",
                    "ProjectName": "PROJECT",
                    "TimeSpanMinutes": 90.0,
                    "TimeSpanQuarterHours": 6.0,
                    "TimeSpanHours": 1.5,
                    "BillableHours": 0.75,
                    "ReportLine": "0900-1030 example notes line 2"
                }
            },
            "Name": "PROJECT",
            "ProjectTotal": 1.75,
            "ClockTime": 2.25,
            "Report": ""
        }
    },
    "Notes": [
        "example notes line",
        "example notes line 2"
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
        },
        {
            "BeginTime": "2014-04-21T09:00:00",
            "EndTime": "2014-04-21T10:30:00",
            "Fraction": 0.5,
            "Notes": "example notes line 2",
            "ProjectName": "PROJECT",
            "TimeSpanMinutes": 90.0,
            "TimeSpanQuarterHours": 6.0,
            "TimeSpanHours": 1.5,
            "BillableHours": 0.75,
            "ReportLine": "0900-1030 example notes line 2"
        }
    ]
}`

func TestImportReportFromJSON(t *testing.T) {
	report, err := ImportReportFromJSON(strings.NewReader(jsonStr))
	if err != nil {
		t.Error(err)
	}

	records, projects := ConvertLegacyRecords(report)
	t.Log(records)
	t.Log(projects)
}
