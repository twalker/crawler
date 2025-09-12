package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := os.Args[1]
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("base url could not be parsed:", err)
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)
	cfg := config{
		baseURL:            baseURL,
		pages:              make(map[string]int),
		concurrencyControl: make(chan struct{}, 1),
		mu:                 &sync.Mutex{},
	}
	cfg.crawlPage(rawBaseURL)

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
