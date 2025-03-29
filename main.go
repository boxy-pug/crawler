package main

import (
	"fmt"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[1]

	//fmt.Printf("starting crawl of: %s\n", baseURL)

	pages := make(map[string]int)

	crawlPage(baseURL, baseURL, pages)

	for key, val := range pages {
		fmt.Printf("Found %d occurrences of page %s\n", val, key)
	}

}
