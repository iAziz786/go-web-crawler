package main

import (
	"fmt"
	"log"

	"github.com/iAziz786/go-web-crawler/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
