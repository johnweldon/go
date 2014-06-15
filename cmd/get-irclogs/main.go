package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var wg = []sync.WaitGroup{sync.WaitGroup{}, sync.WaitGroup{}}
var urls = make(chan downloadData, 4)

type downloadData struct {
	URL  string
	File string
}

type generateArgs struct {
	BaseURL string
	Target  string
	Begin   time.Time
}

// Main is a testable delegate of main
func Main(args []string) {
	begin, _ := time.Parse("20060102", "20100101")
	genArgs := generateArgs{
		BaseURL: "http://irclogs.ubuntu.com",
		Target:  "#juju-dev.txt",
		Begin:   begin,
	}
	wg[0].Add(2)
	go generateDates(genArgs)
	go consumeURLs()
	wg[0].Wait()
}

func generateDates(args generateArgs) {
	defer wg[0].Done()

	date := args.Begin
	now := time.Now()
	for {
		if now.Before(date) {
			close(urls)
			return
		}
		urls <- downloadData{URL: date.Format(args.BaseURL + "/2006/01/02/" + args.Target), File: date.Format("2006-01-02-" + args.Target[1:])}
		date = date.AddDate(0, 0, 1)
		select {
		case <-time.After(100 * time.Millisecond):
			continue
		}
	}
}

func consumeURLs() {
	defer wg[0].Done()

	for {
		select {
		case u, ok := <-urls:
			if ok {
				wg[1].Add(1)
				go downloadURL(u)
			} else {
				wg[1].Wait()
				return
			}
		}
	}
}

func downloadURL(dd downloadData) {
	defer wg[1].Done()

	u, err := url.Parse(url.QueryEscape(dd.URL))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't parse %s: %q\n", dd.URL, err)
		return
	}
	response, err := http.Get(u.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't get %s: %q\n", u, err)
		return
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		fmt.Fprintf(os.Stdout, "Found: %s; saving to %s\n", u, dd.File)

		out, err := os.Create(dd.File)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't create output file %s; %q\n", dd.File, err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't download file %s; %q\n", dd.File, err)
			return
		}
	case http.StatusNotFound:
		return
	default:
		fmt.Fprintf(os.Stdout, "Unexpected response %+v\n", response.Status)
	}
}

func main() {
	flag.Parse()
	Main(flag.Args())
}
