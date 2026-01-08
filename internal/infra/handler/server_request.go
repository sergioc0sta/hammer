package handler

import (
	"net/http"
)

func ServerRequestHandler(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
