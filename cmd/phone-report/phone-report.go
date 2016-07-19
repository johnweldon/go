package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/johnweldon/go-misc/web"
)

func main() {

	flag.Parse()

	for _, num := range flag.Args() {
		printNumber(formatNumber(num))
	}

}

func formatNumber(number string) string {
	re := regexp.MustCompile(".*1\\?.*(\\d{3}).*(\\d{3}).*(\\d{4})")
	res := re.FindStringSubmatch(number)
	if len(res) == 4 {
		return "1-" + strings.Join(res[1:], "-")
	}
	return number
}

func printNumber(number string) {
	fmt.Fprintf(os.Stdout, "Querying for number '%s' ...    ", number)

	scraper := web.NewScraper()
	scraper.ScrapePredicates = []web.ScrapePredicate{web.HasClass("oos_p6")}
	scraper.ScrapeAction = web.TextChildren
	addr := "http://800notes.com/Phone.aspx/" + number
	s, e := scraper.Get(addr)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error: %#v\n", e)
	}
	fmt.Fprintf(os.Stdout, "%d Results:\n-------\n", len(s.Matches()))
	for _, m := range s.Matches() {
		fmt.Fprintf(os.Stdout, "%s\n  ------- ===== -------\n", m)
	}
}
