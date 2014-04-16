package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var stdout, stderr, stddebug *os.File

func main() {
	var root string
	var out, err string
	var debug bool
	var e error
	globber := NewGlobber()
	globber2 := NewGlobber()
	globber2.Default = true

	flag.Var(globber, "X", "exclude pattern")
	flag.Var(globber2, "I", "include pattern")
	flag.StringVar(&root, "r", ".", "root to begin searching")
	flag.StringVar(&out, "o", "", "redirect output to file")
	flag.StringVar(&err, "e", "", "redirect errors to file")
	flag.BoolVar(&debug, "d", false, "debug output")
	flag.Parse()

	stdout = os.Stdout
	if len(out) != 0 {
		stdout, e = os.OpenFile(out, os.O_RDWR|os.O_CREATE, 666)
		if e != nil {
			stdout = os.Stdout
		}
	}
	stderr = os.Stderr
	if len(err) != 0 {
		stderr, e = os.OpenFile(err, os.O_RDWR|os.O_CREATE, 666)
		if e != nil {
			stderr = os.Stderr
		}
	}
	stddebug, e = os.OpenFile(os.DevNull, os.O_RDWR|os.O_CREATE, 666)
	if debug || e != nil {
		stddebug = stderr
	}

	args := flag.Args()
	for _, val := range args {
		g := NewGrepVisitor(val)
		g.Exclude = globber
		g.Include = globber2
		grep(root, g)
	}
}

func grep(root string, v *GrepVisitor) { filepath.Walk(root, v.Visit) }

type Globber struct {
	Pattern map[int]string
	Default bool
}

func NewGlobber() (g *Globber) {
	g = new(Globber)
	g.Pattern = make(map[int]string)
	g.Default = false
	return g
}

func (g *Globber) String() string {
	return "globber"
}

func (g *Globber) Set(str string) error {
	for _, p := range strings.Split(str, ",") {
		g.Pattern[len(g.Pattern)] = p
	}
	return nil
}

func (g *Globber) Match(filepath string) bool {
	fmt.Fprintf(stddebug, "Does '%s' match '%v'?...  ", filepath, g.Pattern)
	if len(g.Pattern) == 0 {
		fmt.Fprintf(stddebug, "%v (default)\n", g.Default)
		return g.Default
	}

	_, file := path.Split(filepath)
	for _, pattern := range g.Pattern {
		m, e := path.Match(pattern, file)
		if e == nil && m {
			fmt.Fprintf(stddebug, "yes\n")
			return true
		}
	}
	fmt.Fprintf(stddebug, "no\n")
	return false
}

type GrepVisitor struct {
	Pattern   *regexp.Regexp
	MatchOnly bool
	Exclude   *Globber
	Include   *Globber
}

func NewGrepVisitor(re string) (v *GrepVisitor) {
	v = new(GrepVisitor)
	v.MatchOnly = true
	v.Exclude = NewGlobber()
	v.Include = NewGlobber()
	v.Include.Default = true

	rex, err := regexp.Compile(re)
	if err != nil {
		fmt.Fprintf(stderr, "COMPILE ERROR: %+v\n", err)
		return v
	}
	v.Pattern = rex

	return v
}

func (v *GrepVisitor) Visit(filepath string, f os.FileInfo, err error) error {
	return nil
}

func (v *GrepVisitor) VisitDir(filepath string, f os.FileInfo) bool {
	if v.Pattern == nil {
		return false
	}
	if v.Exclude.Match(filepath) {
		return false
	}
	if v.Include.Match(filepath) {
		return true
	}
	return true
}

func (v *GrepVisitor) VisitFile(filepath string, f os.FileInfo) {
	if v.Pattern == nil {
		return
	}
	if v.Exclude.Match(filepath) {
		return
	}
	if v.Include.Match(filepath) {
		processFile(filepath, v)
	}

}

func processFile(filepath string, v *GrepVisitor) {
	raw, e1 := ioutil.ReadFile(filepath)
	if e1 != nil {
		fmt.Fprintf(stderr, "IOERROR: %+v\n", e1)
		return
	}
	if v.Pattern.Match(raw) {
		fmt.Fprintf(stdout, "%s\n", filepath)
	}
}
