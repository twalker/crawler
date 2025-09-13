package main

import (
	"fmt"
	"os"
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
	// baseURL, err := url.Parse(rawBaseURL)
	// if err != nil {
	// 	fmt.Println("base url could not be parsed:", err)
	// }

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	//cfg.crawlPage(rawBaseURL)

	cfg := NewCrawler(rawBaseURL, 5)
	//cfg.wg.Add(1)
	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
