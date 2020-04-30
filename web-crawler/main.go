/*
	В этом упражнении вы будете использовать возможности
	многопоточности Go для распараллеливания поискового робота.
	Измените функцию Crawl для извлечения URL параллельно,
	не запрашивая один и тот же URL дважды.
*/

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SaveMap struct {
	m   map[string]bool
	mux sync.Mutex
}

func (s *SaveMap) IsItem(url string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.m[url]
	return ok
}

func (s *SaveMap) Insert(url string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[url] = true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, caughtUrl *SaveMap) {
	if depth <= 0 || caughtUrl.IsItem(url) {
		return
	}
	caughtUrl.Insert(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("depth: %v | found: %s %q\n", depth, url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, caughtUrl)
	}
}

func main() {
	caughtUrl := &SaveMap{m: make(map[string]bool)}
	Crawl("http://golang.org/", 4, fetcher, caughtUrl)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
