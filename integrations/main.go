package main

import (
	"net/http"
	"os"
	"time"
)

// Used by integrations-entrypoint.sh to check if the entrypoint is up
func main() {
	// URL of the service we're waiting for is passed via CLI arg
	url := os.Args[1]

	http_client := &http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		os.Exit(1)
	}

	resp, err := http_client.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	// form3tech/interview-accountapi image returns 404 for an index GET
	if resp.StatusCode != 404 {
		os.Exit(1)
	}

	os.Exit(0)
}
