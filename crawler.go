package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if !isSameDomain(rawBaseURL, rawCurrentURL) {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawBaseURL, rawCurrentURL)
	if err != nil {
		fmt.Printf("Could not normalize %v: %v\n", rawCurrentURL, err)
		return
	}

	_, ok := pages[normalizedCurrentURL]
	if ok {
		pages[normalizedCurrentURL]++
		return
	}
	pages[normalizedCurrentURL] = 1

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
		fmt.Printf("crawling %v\n", url)
		crawlPage(normalizedURLWithHTTPS, url, pages)
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
