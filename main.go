package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {

	maxConcurrency := 5
	maxPages := 10

	args := os.Args

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		var err error
		maxConcurrency, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("could not parse %v as int for max concurrency\n", args[2])
			os.Exit(1)
		}
		fmt.Printf("max concurrency set to %v\n", maxConcurrency)
	}
	if len(args) > 3 {
		var err error
		maxPages, err = strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("could not parse %v as int for max pages\n", args[3])
			os.Exit(1)
		}
		fmt.Printf("max pages set to %v\n", maxPages)
	}

	cfg := &config{
		baseURL:            args[1],
		pages:              make(map[string]int),
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	//fmt.Printf("starting crawl of: %s\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL)

	cfg.wg.Wait()

	for key, val := range cfg.pages {
		fmt.Printf("Found %d occurrences of page %s\n", val, key)
	}

}
