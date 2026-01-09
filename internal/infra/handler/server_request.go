package handler

import (
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func ServerRequestHandler(url string) (int, error) {
	resp, err := httpClient.Get(url)

	if err != nil {
		return 1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
