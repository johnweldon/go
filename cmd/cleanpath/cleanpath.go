package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	clean := cleanPath(args)
	fmt.Fprintf(os.Stdout, "%s", clean)
}

func cleanPath(paths []string) string {
	pos := 0
	working := make(map[string]int)
	for _, pth := range paths {
		for _, segment := range strings.FieldsFunc(pth, func(c rune) bool { return c == os.PathListSeparator }) {
			clean := path.Clean(segment)
			// ignore relative paths
			if !path.IsAbs(clean) {
				continue
			}
			// Make sure relative order remains, and if duplicates then keep the first position
			if _, present := working[clean]; !present {
				working[clean] = pos
				pos++
			}
		}
	}
	array := make([]string, len(working))
	for segment, position := range working {
		array[position] = segment
	}
	return strings.Join(array, string(os.PathListSeparator))
}
