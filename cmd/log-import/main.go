package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"

	logf "github.com/johnweldon/go/log"
)

var importPath string

func init() {
	flag.StringVar(&importPath, "import", "", "path to the json formatted import file")
}

func main() {
	flag.Parse()
	reader, err := os.Open(importPath)
	if err != nil {
		log.Fatal(err)
	}

	report, err := logf.ImportReportFromJson(reader)
	if err != nil {
		log.Fatal(err)
	}

	records, projects := logf.ConvertLegacyRecords(report)

	b, err := json.Marshal(struct {
		Records  []logf.TimeRecord
		Projects []logf.Project
	}{records, projects})
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("imported.json")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	buf.WriteTo(out)
}
