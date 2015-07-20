package web

import (
	"net/http"
	"strings"

	"code.google.com/p/go.net/html"
)

type ScrapeResult interface {
	Raw() string
	Matches() []interface{}
}

type Scraper interface {
	Get(uri string) (shavings ScrapeResult, err error)
}

type ScrapePredicate func(*html.Node) bool
type ScrapeAction func(*html.Node) (string, error)

type Scrape struct {
	ScrapePredicates []ScrapePredicate
	ScrapeAction     ScrapeAction
}

func NewScraper() *Scrape {
	return &Scrape{
		ScrapePredicates: []ScrapePredicate{
			defaultWebScrapePredicate,
		},
		ScrapeAction: defaultWebScrapeAction,
	}
}

func (ws *Scrape) Get(uri string) (shavings ScrapeResult, err error) {
	result := &webScrapeResult{}
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	body, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	var scrape func(*html.Node)
	scrape = func(n *html.Node) {
		if any(ws.ScrapePredicates, n) {
			match, err := ws.ScrapeAction(n)
			if err == nil {
				result.matches = append(result.matches, match)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			scrape(c)
		}
	}
	scrape(body)
	result.raw = strings.Join(result.matches, "\r\n")
	shavings = result
	return shavings, nil
}

func HasClass(class string) ScrapePredicate {
	return HasAttrVal("class", class)
}

func HasAttrVal(name, value string) ScrapePredicate {
	return func(n *html.Node) bool {
		pass := false
		attributeAction(n,
			func(a html.Attribute) bool { return a.Key == name },
			func(a html.Attribute) {
				for _, val := range strings.Split(a.Val, " ") {
					if val == value {
						pass = true
					}
				}
			})
		return pass
	}
}

func HasAttr(name string) ScrapePredicate {
	return func(n *html.Node) bool {
		pass := false
		attributeAction(n,
			func(a html.Attribute) bool { return !pass && a.Key == name },
			func(a html.Attribute) { pass = true })
		return pass
	}
}

func TextChildren(n *html.Node) (string, error) {
	if n == nil {
		return "", nil
	}
	var ret string
	if n.Type == html.TextNode {
		ret = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			ret += "\r\n" + c.Data
		}
	}
	return ret, nil
}

func AttrValue(name string) ScrapeAction {
	return func(n *html.Node) (string, error) {
		var r string
		attributeAction(n,
			func(a html.Attribute) bool { return a.Key == name },
			func(a html.Attribute) { r = a.Val })
		return r, nil
	}
}

func defaultWebScrapeAction(n *html.Node) (string, error) {
	return n.Data, nil
}

func defaultWebScrapePredicate(n *html.Node) bool {
	if n == nil {
		return false
	}
	return n.Type == html.TextNode
}

type webScrapeResult struct {
	raw     string
	matches []string
}

func (wsr *webScrapeResult) Raw() string {
	return wsr.raw
}
func (wsr *webScrapeResult) Matches() []interface{} {
	ret := make([]interface{}, len(wsr.matches))
	for ix, v := range wsr.matches {
		ret[ix] = v
	}
	return ret
}

func any(predicates []ScrapePredicate, match *html.Node) bool {
	for _, fn := range predicates {
		if fn(match) {
			return true
		}
	}
	return false
}

func attributeAction(n *html.Node, pred func(html.Attribute) bool, fn func(html.Attribute)) {
	if n != nil {
		for _, attr := range n.Attr {
			if pred(attr) {
				fn(attr)
			}
		}
	}
}
