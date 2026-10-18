package main

import (
	"fmt"
	"os"
)

func main() {
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
}
