package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	} else {
		fmt.Println("current URL", rawCurrentURL)
	}
	// skip other websites
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	// normalizedURL, err := normalizeURL(rawCurrentURL)
	// if err != nil {
	// 	fmt.Printf("Error - normalizedURL: %v", err)
	// 	return
	// }
	normalizedURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v", err)
		return
	}
	isFirstVisit := cfg.addPageVisit(normalizedURL.String())
	if !isFirstVisit {
		return
	}
	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v", err)
		return
	}

	for _, nextURL := range nextURLs {
		fmt.Println(nextURL)
		cfg.crawlPage(nextURL)
	}
}
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	// mark as visited
	cfg.pages[normalizedURL] = 1
	return true
}
