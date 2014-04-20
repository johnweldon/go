package web

import (
	"net"
	"net/http"

	"testing"
)

var testResponse string = "<http><head></head><body><div>Raw Text</div><p class='test3'>Paragraph</p><span>span</span></body></http>"

type testServerHandler struct{}

func (h *testServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(testResponse))
}

func TestScrape(t *testing.T) {
	l, e := startServer()
	if e != nil {
		t.Error(e)
	}
	defer l.Close()

	scraper := NewWebScraper()
	addr := "http://" + l.Addr().String() + "/blah/blah"
	t.Logf("Addr '%s'\n", addr)
	s, e := scraper.Scrape(addr)
	if e != nil {
		t.Error(e)
	}
	if s == nil {
		t.Error("Returned nothing")
	}
	if s.Raw() == "" {
		t.Error("Raw result is empty")
	}
	if len(s.Matches()) != 3 {
		t.Error("Found wrong number of matches", len(s.Matches()), s.Matches())
	}
}

func TestCustomScrape(t *testing.T) {
	l, e := startServer()
	if e != nil {
		t.Error(e)
	}
	defer l.Close()

	scraper := NewWebScraper()
	scraper.ScrapePredicates = []ScrapePredicate{HasClass("test3")}
	scraper.ScrapeAction = TextChildren
	addr := "http://" + l.Addr().String() + "/blah/blah"
	t.Logf("Addr '%s'\n", addr)
	s, e := scraper.Scrape(addr)

	if e != nil {
		t.Error(e)
	}
	if s == nil {
		t.Error("Returned nothing")
	}
	if s.Raw() == "" {
		t.Error("Raw result is empty")
	}
	if len(s.Matches()) != 1 {
		t.Error("Found wrong number of matches", len(s.Matches()), s.Matches())
	}
}

func startServer() (net.Listener, error) {
	l, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		return nil, e
	}
	go http.Serve(l, &testServerHandler{})
	return l, nil
}
