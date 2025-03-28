package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	strippedPath := strings.TrimRight(parsedURL.Path, "/")
	cleanURL := fmt.Sprintf("%v%v", parsedURL.Host, strippedPath)

	return cleanURL, nil
}
