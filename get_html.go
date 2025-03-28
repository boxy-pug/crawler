package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 399 {
		return "", fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return "", err
	}
	contentHeaders := res.Header.Get("Content-Type")

	if strings.HasPrefix(contentHeaders, "text/html") {
		return string(body), nil
	}

	return "", fmt.Errorf("content type header not text/html, %v instead", res.Header.Get("content-type header"))

}

/*
Use http.Get to fetch the webpage of the rawURL
Return an error if the HTTP status code is an error-level code (400+)
Return an error if the response content-type header is not text/html
Return any other possible errors
Return the webpage's HTML if successful

You may find io.ReadAll helpful in reading the response.
*/
