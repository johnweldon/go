package main

import (
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

	log.Printf("Projects (%d)\n%s\n", len(report.Projects), report.Projects)
}
