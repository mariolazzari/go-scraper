package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func getHTML(rawURL string) (string, error) {

	// request
	req, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return "", err
	}
	// set request header to avoid blocks
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	// client
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	// make request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// check response status
	if res.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("bad request: %d - %s", res.StatusCode, res.Status)
	}

	// check response content type
	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("bad content type: %s", contentType)
	}

	// read body
	reader, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body response: %s", err)
	}

	return string(reader), nil
}
