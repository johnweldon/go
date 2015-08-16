package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Create("dump.bin")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	defer f.Close()

	var nr nullReader
	w, err := io.CopyN(f, nr, 20*1024*1024)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	fmt.Fprintf(os.Stdout, "wrote %d bytes\n", w)
}

type nullReader struct{}

func (r nullReader) Read(p []byte) (int, error) {
	for i, _ := range p {
		p[i] = 0
	}
	return len(p), nil
}
