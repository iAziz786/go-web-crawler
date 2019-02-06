package main

import (
	"net/url"
	"os"
)

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	// Number of pending sends to the worklist
	const MaxCralwer = 20

	crawlableURLs := os.Args[1:]
	n := len(crawlableURLs)
	go func() { worklist <- os.Args[1:] }()

	// Start 20 goroutines, each of them will wait for a link to be sent down the
	// `unseenLinks` channel. After crawl they will push data into the worklist channel
	for i := 0; i < MaxCralwer; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if didHostnameMatch(crawlableURLs, link) {
				if !seen[link] {
					seen[link] = true
					n++
					unseenLinks <- link
				}
			}
		}
	}
}

func didHostnameMatch(urls []string, link string) bool {
	parsedCandidatedURL, _ := url.Parse(link)
	for _, eachURL := range urls {
		parsedURL, _ := url.Parse(eachURL)
		if parsedURL.Hostname() == parsedCandidatedURL.Hostname() {
			return true
		}
	}
	return false
}
