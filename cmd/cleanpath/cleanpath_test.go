package main

import (
	"testing"
)

func TestCleanpath(t *testing.T) {
	for input, expected := range testCase1 {
		got := cleanPath([]string{input})
		if expected != got {
			t.Error("Expected '" + expected + "', got '" + got + "'\n")
		}
	}
}
