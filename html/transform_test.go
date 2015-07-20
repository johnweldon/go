package html

import (
	"encoding/hex"
	"strings"
	"testing"
)

var testDocs1 = map[string]string{
	"\n\n<html><body><div><lI clAss='gg aA ea'>HELLO</li><li>IGNORE</li><li class='gg'>MORE IGNORE</li>": "HELLO",
	"<div><li class='aa'>HELLO</li><li>IGNORE</li><li class='aa'>KITTY</li>":                             "HELLO\nKITTY",
}

var testDocs2 = map[string]string{
	"<div></p><a href='about:blank'></a>":                                                                         "about:blank",
	"<html><head></head><div><a href='about:test'></a><a src='nil'/><a href='about:test2'>test2</a></div></body>": "about:test\nabout:test2",
}

func TestTransformGetText(t *testing.T) {
	x := NewTransformer(ElementWithClass("li", "aa"), GetAllText())
	for input, expected := range testDocs1 {
		output, err := x.Transform(strings.NewReader(input))
		if err != nil {
			t.Error(err)
		}
		if output != expected {
			t.Error("\nExpected: x'" + expected + "'(" + hex.EncodeToString([]byte(expected)) + ")\n     Got: x'" + output + "'(" + hex.EncodeToString([]byte(output)) + ")")
		}
	}
}

func TestTransformGetLinks(t *testing.T) {
	x := NewTransformer([]Predicate{Element("div")}, GetAllLinks())
	for input, expected := range testDocs2 {
		output, err := x.Transform(strings.NewReader(input))
		if err != nil {
			t.Error(err)
		}
		if output != expected {
			t.Error("\nExpected: x'" + hex.EncodeToString([]byte(expected)) + "'\n     Got: x'" + hex.EncodeToString([]byte(output)) + "'")
		}
	}
}
