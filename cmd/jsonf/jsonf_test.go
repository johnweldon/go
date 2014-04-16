package main

import (
	"bytes"
	"io"
	"testing"
)

func TestMain(t *testing.T) {
	var reader io.Reader
	var writer io.Writer

	data := []byte(`{"Name":"John","Active":"true"}`)
	reader = bytes.NewBuffer(data)
	writer = new(bytes.Buffer)

	if err := Format(reader, writer); err != nil {
		t.Errorf("Error: %v [buffer: %v]", err, writer)
	}
	t.Logf("%s\n", writer)
}
