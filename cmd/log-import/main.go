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
var exportPath string
var serverlist string

func init() {
	flag.StringVar(&importPath, "import", "", "path to the json formatted legacy format input file")
	flag.StringVar(&exportPath, "export", "imported.json", "path to the json formatted output file")
	flag.StringVar(&serverlist, "serverlist", "localhost", "mongo serverlist")
}

func main() {
	flag.Parse()
	records, projects := convert()
	writeFile(records, projects)
	writeMongo(records)
}

func writeMongo(records []logf.TimeRecord) {
	db := logf.NewDB(serverlist)
	err := db.SaveRecords(records)
	if err != nil {
		panic(err)
	}
	readin := db.GetRecords()
	log.Printf("Imported %d records\n", len(readin))
}

func convert() ([]logf.TimeRecord, []logf.Project) {
	reader, err := os.Open(importPath)
	if err != nil {
		log.Fatal(err)
	}

	report, err := logf.ImportReportFromJson(reader)
	if err != nil {
		log.Fatal(err)
	}

	return logf.ConvertLegacyRecords(report)
}

func writeFile(records []logf.TimeRecord, projects []logf.Project) {
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

	out, err := os.Create(exportPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	buf.WriteTo(out)
}
