/*
Sanitize

I use this in git clean/smudge filters to filter files that I don't want to check in as is.

For example:

In this example my working copy of the file `my_private_file.config` has my real credentials.

However when I put this sanitize tool on my path and modify my git config like this it causes the
sensitive data in `my_private_file.config` to be replaced per my configuration in `.git/config` before
being added to the repository.


my_private_file.config:
    username=john.weldon
    password=s3krit



.gitattributes file:
    my_private_file.config      filter=sanitize



.git/config file:
    [filter "sanitize"]
        clean = sanitize "/john.weldon/sample.username/" "/s3krit/sample.password/"
        smudge = sanitize "/sample.username/john.weldon/" "/sample.password/s3krit/"
        required

*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	process(getReplacements(flag.Args()), os.Stdin, os.Stdout)
}

func process(replacements []string, inStream io.Reader, outStream io.Writer) {
	in := bufio.NewScanner(inStream)
	out := bufio.NewWriter(outStream)
	replacer := strings.NewReplacer(replacements...)
	for in.Scan() {
		fmt.Fprintln(out, replacer.Replace(in.Text()))
	}
	out.Flush()
}

func isPattern(arg string) bool {
	return len(arg) > 3 &&
		string(arg[0]) == string(arg[len(arg)-1]) &&
		strings.Count(arg, string(arg[0])) == 3
}

func getReplacements(args []string) []string {
	r := make([]string, len(args)*2)
	for _, arg := range args {
		if isPattern(arg) {
			sp := strings.Split(arg, string(arg[0]))
			r = append(r, sp[1])
			r = append(r, sp[2])
		}
	}
	return r
}
