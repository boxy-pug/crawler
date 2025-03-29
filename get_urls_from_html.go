package main

import (
	"bytes"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var res []string

	//z := html.NewTokenizer(r)

	doc, err := html.Parse(bytes.NewReader([]byte(htmlBody)))
	if err != nil {
		return res, err
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					res = append(res, attr.Val)
				}
			}
		}
	}

	return res, nil
}
