package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var reader io.Reader
	var writer io.Writer
	reader = os.Stdin
	writer = os.Stdout
	Format(reader, writer)
}

func Format(reader io.Reader, writer io.Writer) error {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("Error reading source: %v", err)
	}
	dst := new(bytes.Buffer)
	if err = json.Indent(dst, buf, "", "  "); err != nil {
		return fmt.Errorf("Error indenting: %v", err)
	}
	fmt.Fprintf(writer, "%s\n", dst)
	return nil
}
