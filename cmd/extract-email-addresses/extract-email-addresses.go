package main

import (
	"flag"
	"fmt"
	"io"
	"net/mail"
	"os"
)

var (
	CRLF  = []byte("\r\n")
	BEGIN = []byte("BEGIN:VCARD\r\nVERSION:4.0\r\n")
	END   = []byte("END:VCARD\r\n")
)

var inFile string
var outFile string

func init() {
	flag.StringVar(&inFile, "in", "in.eml", "name of email file")
	flag.StringVar(&outFile, "out", "out.vcf", "name of vcf file to put extracted email addresses")
}

func main() {

	flag.Parse()

	var vcf *os.File
	var err error
	var r io.Reader

	var msg *mail.Message

	if r, err = os.Open(inFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v\n", err)
		return
	}

	if msg, err = readMessage(r); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v\n", err)
		return
	}

	if vcf, err = os.Create(outFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v\n", err)
		return
	}
	defer vcf.Close()

	writeMessageVCF(vcf, msg)

}

func readMessage(reader io.Reader) (*mail.Message, error) {
	return mail.ReadMessage(reader)
}

func writeMessageVCF(writer io.Writer, message *mail.Message) {
	var err error
	var list []*mail.Address
	for _, header := range []string{"To", "Cc", "Bcc", "From"} {
		if list, err = message.Header.AddressList(header); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading header %s: %#v\n", header, err)
			continue
		}
		for _, address := range list {
			writeAddressVCF(writer, address)
		}
	}
}

func writeAddressVCF(writer io.Writer, address *mail.Address) {
	writer.Write(BEGIN)
	writer.Write([]byte("FN:" + address.Name))
	writer.Write(CRLF)
	writer.Write([]byte("EMAIL;TYPE=INTERNET:" + address.Address))
	writer.Write(CRLF)
	writer.Write(END)
}
