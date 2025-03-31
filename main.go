package main

import (
	"fmt"
	"os"
	"sort"
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

type urlCount struct {
	URL   string
	Count int
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

	cfg.printReport()

	/*
		for key, val := range cfg.pages {
			fmt.Printf("Found %d occurrences of page %s\n", val, key)
		}
	*/

}

func (cfg *config) printReport() {
	reportBorder := "============================="
	fmt.Println(reportBorder)
	fmt.Printf("REPORT for %s\n", cfg.baseURL)
	fmt.Println(reportBorder)

	sortedPages := sortPagesByVal(cfg.pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to https://%s\n", page.Count, page.URL)
	}
}

func sortPagesByVal(pages map[string]int) []urlCount {
	urlCounts := make([]urlCount, 0, len(pages))

	for k, v := range pages {
		urlCounts = append(urlCounts, urlCount{URL: k, Count: v})
	}

	// Sort with custom comparator
	sort.Slice(urlCounts, func(i, j int) bool {
		// First compare by count (higher first)
		if urlCounts[i].Count != urlCounts[j].Count {
			return urlCounts[i].Count > urlCounts[j].Count
		}
		// If counts are equal, sort alphabetically
		return urlCounts[i].URL < urlCounts[j].URL
	})

	return urlCounts

}
