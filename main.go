package main

import (
	"net/url"
	"os"
)

func main() {
	worklist := make(chan []string)
	// Number of pending sends to the worklist
	var n int

	n++
	crawlableURLs := os.Args[1:]
	go func() { worklist <- os.Args[1:] }()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if didHostnameMatched(crawlableURLs, link) {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
	}
}

func didHostnameMatched(urls []string, link string) bool {
	parsedCandidatedURL, _ := url.Parse(link)
	for _, eachURL := range urls {
		parsedURL, _ := url.Parse(eachURL)
		if parsedURL.Hostname() == parsedCandidatedURL.Hostname() {
			return true
		}
	}
	return false
}
