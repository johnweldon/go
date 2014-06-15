package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"code.google.com/p/go.net/html"
)

func Main(args []string) {
	baseurl := "http://irclogs.ubuntu.com"
	urls := []string{baseurl}
	target := "#juju-dev.txt"

	client := http.Client{}
	for {
		if len(urls) == 0 {
			break
		}
		cur := urls[0]
		urls = urls[1:]
		prefix := strings.Replace(cur[len(baseurl):], "/", "-", 0)

		fmt.Fprintf(os.Stderr, "checking %s\n", cur)
		response, err := client.Get(cur)
		if err != nil {
			fmt.Fprintf(os.Stderr, "problem reading '%s'; %q\n", cur, err)
			continue
		}
		matches, links := find(response, baseurl, target)
		urls = append(urls, links...)
		for _, match := range matches {
			saveMatch(&client, prefix, target, match)
		}
	}

}

func saveMatch(c *http.Client, prefix, target, match string) {
	fmt.Fprintf(os.Stdout, "MATCH: %s_%s\n", prefix, target[1:])
}

func find(r *http.Response, baseurl, target string) (matches, links []string) {
	if r.StatusCode == http.StatusOK {
		base := r.Request.URL.String()
		baseUrl, err := r.Location()
		if err == nil {
			base = baseUrl.String()
		}
		body, err := html.Parse(r.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "html parse error %q\n", err)
			return
		}

		var findFn func(*html.Node)
		findFn = func(n *html.Node) {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, "?") {
						continue
					}
					if strings.HasPrefix(attr.Val, "#") {
						continue
					}
					if strings.HasPrefix(attr.Val, "http") && !strings.HasPrefix(attr.Val, baseurl) {
						continue
					}
					link := base + "/" + attr.Val
					if strings.HasSuffix(link, target) {
						matches = append(matches, link)
					} else {
						links = append(links, link)
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				findFn(c)
			}
		}

		findFn(body)
	}
	return
}

func main() {
	flag.Parse()
	Main(flag.Args())
}
