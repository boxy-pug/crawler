package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
	}()

	if !isSameDomain(cfg.baseURL, rawCurrentURL) {
		return
	}

	normalizedCurrentURL, err := normalizeURL(cfg.baseURL, rawCurrentURL)
	if err != nil {
		fmt.Printf("Could not normalize %v: %v\n", rawCurrentURL, err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	normalizedURLWithHTTPS := fmt.Sprintf("https://%s", normalizedCurrentURL)

	currentHTML, err := getHTML(normalizedURLWithHTTPS)
	if err != nil {
		fmt.Printf("could not get html from %v: %v\n", normalizedURLWithHTTPS, err)
		return
	}
	fmt.Printf("fetched html from %v\n", normalizedURLWithHTTPS)

	urlsFromHTML, err := getURLsFromHTML(currentHTML, normalizedURLWithHTTPS)
	if err != nil {
		fmt.Printf("could not get urls from %v: %v\n", normalizedURLWithHTTPS, err)

	}

	for _, url := range urlsFromHTML {
		if cfg.isMaxPagesReached() {
			break
		}
		fmt.Printf("scheduling crawl for %v\n", url)

		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}

}

func isSameDomain(base, url2 string) bool {

	baseURL, err := url.Parse(base)
	if err != nil {
		return false
	}
	parsedURL2, err := url.Parse(url2)
	if err != nil {
		return false
	}

	resolvedURL2 := baseURL.ResolveReference(parsedURL2)

	return baseURL.Host == resolvedURL2.Host

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if len(cfg.pages) >= cfg.maxPages {
		return false
	}

	_, ok := cfg.pages[normalizedURL]
	if ok {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true

}

func (cfg *config) isMaxPagesReached() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages) >= cfg.maxPages

}
