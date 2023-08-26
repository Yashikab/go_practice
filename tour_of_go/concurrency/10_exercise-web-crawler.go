package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth
func Crawl(url string, depth int, fetcher Fetcher, ch0 chan int) {
	// TODO: Fetch URLs in parallel
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	var mutex sync.Mutex
	// defer func() {
	// 	ch0 <- 0
	// }()
	if depth <= 0 {
		ch0 <- 0
		return
	}
	_, ok := fetchedURL[url]
	if ok {
		ch0 <- 0
		return
	} else {
		mutex.Lock()
		fetchedURL[url] = true
		mutex.Unlock()
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch0 <- 0
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch1 := make(chan int)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, ch1)
	}
	for i := 0; i < len(urls); i++ {
		<-ch1
	}
	ch0 <- 0
	return
}

func main() {
	ch := make(chan int)
	Crawl("https://golang.org/", 4, fetcher, ch)
	<-ch
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetchedURL = make(map[string]bool)

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
