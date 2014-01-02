package main

import (
	"fmt"
	stat "github.com/johnweldon/go/go-stat"
	"os"
)

func main() {
	stat.Open()
	fmt.Fprintf(os.Stdout, "Hello\n")
}
