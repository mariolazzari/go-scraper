package main

import (
	"fmt"
	"os"
)

func main() {
	// check args
	argsLen := len(os.Args[1:])
	if argsLen < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if argsLen > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseUrl)

	// read html
	pages := make(map[string]int)

	crawlPage(baseUrl, baseUrl, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
