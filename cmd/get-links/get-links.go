package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/johnweldon/go/web"
)

type Links struct {
	Root  string
	Links []string
}

func (l Links) String() string {
	r := fmt.Sprintf("> %s\n", l.Root)
	for _, v := range l.Links {
		r = fmt.Sprintf("%s- %s\n", r, v)
	}
	return r
}

type S struct {
	web.Scrape
}

func NewS() *S {
	scraper := &S{}
	scraper.ScrapePredicates = []web.ScrapePredicate{web.HasAttr("href")}
	scraper.ScrapeAction = web.AttrValue("href")
	return scraper
}

func (s *S) Crawl(url string) (Links, error) {
	links := Links{url, []string{}}
	r, e := s.Get(url)
	if e != nil {
		return Links{}, e
	}
	for _, v := range r.Matches() {
		links.Links = append(links.Links, v.(string))
	}
	return links, nil
}

func init() {
}

func main() {
	flag.Parse()
	s := NewS()
	for _, root := range flag.Args() {
		r, e := s.Crawl(root)
		if e != nil {
			fmt.Fprintf(os.Stderr, "Error crawling '%s': '%v'\n", root, e)
		}
		fmt.Fprintf(os.Stdout, "%s\n", r)
	}
}
