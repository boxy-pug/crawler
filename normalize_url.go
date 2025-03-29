package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(baseURL, inputURL string) (string, error) {
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %v", err)
	}

	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	resolvedURL := parsedBase.ResolveReference(parsedURL)
	resolvedURL.Path = strings.TrimRight(resolvedURL.Path, "/") // Strip trailing slash

	res := fmt.Sprintf("%s%s", resolvedURL.Host, resolvedURL.Path)

	return res, nil
}
