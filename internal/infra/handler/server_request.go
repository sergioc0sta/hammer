package handler

import (
	"net/http"
)

func ServerRequestHandler(url string, httpClient *http.Client) (int, error) {
	resp, err := httpClient.Get(url)

	if err != nil {
		return 1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
