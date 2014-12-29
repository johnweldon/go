package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"

	jlog "github.com/johnweldon/go/log"
)

var (
	importPath string
	exportPath string
	serverlist string
	csv        bool
)

func init() {
	flag.StringVar(&importPath, "import", "", "path to the json formatted legacy format input file")
	flag.StringVar(&exportPath, "export", "imported.json", "path to the json formatted output file")
	flag.StringVar(&serverlist, "serverlist", "localhost", "mongo serverlist")
	flag.BoolVar(&csv, "csv", false, "expect csv input rather than json")
}

func main() {
	flag.Parse()
	records, projects := convert()
	writeFile(records, projects)
	writeMongo(records)
}

func writeMongo(records []jlog.TimeRecord) {
	db := jlog.NewMongoDB(serverlist)
	err := db.SaveRecords(records)
	if err != nil {
		panic(err)
	}
	readin := db.GetRecords()
	log.Printf("Imported %d records\n", len(readin))
}

func convert() ([]jlog.TimeRecord, []jlog.Project) {
	reader, err := os.Open(importPath)
	if err != nil {
		log.Fatal(err)
	}

	if csv {
		records, err := jlog.ConvertCSVRecords(reader, "")
		if err != nil {
			log.Fatal(err)
		}
		return records, []jlog.Project{jlog.Project{}}
	} else {
		report, err := jlog.ImportReportFromJson(reader)
		if err != nil {
			log.Fatal(err)
		}

		return jlog.ConvertLegacyRecords(report)
	}
}

func writeFile(records []jlog.TimeRecord, projects []jlog.Project) {
	b, err := json.Marshal(struct {
		Records  []jlog.TimeRecord
		Projects []jlog.Project
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
