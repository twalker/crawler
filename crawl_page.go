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

func NewCrawler(rawBaseURL string, maxConcurrency int) *config {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("base url could not be parsed:", err)
	}

	return &config{
		baseURL:            baseURL,
		pages:              make(map[string]int),
		concurrencyControl: make(chan struct{}, maxConcurrency),
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
	}
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		//cfg.wg.Done()
	}()
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
		//cfg.wg.Add(1)
		//go cfg.crawlPage(nextURL)
		cfg.wg.Go(func() { cfg.crawlPage(nextURL) })
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
