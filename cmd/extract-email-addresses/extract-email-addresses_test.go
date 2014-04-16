package main

import (
	"bytes"
	"encoding/hex"
	"net/mail"
	"strings"

	"testing"
)

var message string = "From: test@example.com\r\nTo: fred@example.com\r\nSubject: Test Message\r\n\r\nBody message\r\n"
var vcf string = "BEGIN:VCARD\r\nVERSION:4.0\r\nFN:\r\nEMAIL;TYPE=INTERNET:fred@example.com\r\nEND:VCARD\r\nBEGIN:VCARD\r\nVERSION:4.0\r\nFN:\r\nEMAIL;TYPE=INTERNET:test@example.com\r\nEND:VCARD\r\n"

func TestExtractMessage(t *testing.T) {
	msg := getMessage(t)

	if msg == nil {
		t.Error(msg)
	}
	if msg.Header == nil {
		t.Error(msg)
	}
	if msg.Body == nil {
		t.Error(msg)
	}
}

func TestWriteVCF(t *testing.T) {
	msg := getMessage(t)

	writer := bytes.NewBufferString("")
	writeMessageVCF(writer, msg)

	expected := hex.Dump([]byte(vcf[0:len(vcf)]))
	got := hex.Dump([]byte(writer.String()))
	if expected != got {
		t.Error("Expected\n", expected, "\nGot\n", got)
	}
}

func getMessage(t *testing.T) *mail.Message {
	r := strings.NewReader(message)
	msg, err := readMessage(r)
	if err != nil {
		t.Error(err)
	}
	return msg
}
