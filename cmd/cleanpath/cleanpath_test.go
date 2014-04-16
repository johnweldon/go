package main

import (
	"testing"
)

var testCase1 map[string]string = map[string]string{
	"/a/b/c:/a/b/c:/a/b/c/": "/a/b/c",
	"../a/b:/a/b:../a/b":    "/a/b",
	"/a:/b:/c:/b:/a/:a":     "/a:/b:/c",
}

func TestCleanpath(t *testing.T) {
	for input, expected := range testCase1 {
		got := cleanPath([]string{input})
		if expected != got {
			t.Error("Expected '" + expected + "', got '" + got + "'\n")
		}
	}
}
