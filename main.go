package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("no website provided")
	}
	if len(args) > 1 {
		log.Fatal("too many arguments provided")
	}
	rawBaseUrl := args[0]
	fmt.Printf("starting crawl of: %s\n", rawBaseUrl)
	pages := map[string]int{}
	crawlPage(rawBaseUrl, rawBaseUrl, pages)
	fmt.Println("Results:")
	for k, v := range pages {
		fmt.Printf("\t%s: %d\n", k, v)
	}
}
