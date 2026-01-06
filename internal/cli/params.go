package cli

import (
	"flag"
)

type Params struct {
	URL          string
	Requests     int
	Concurrencies int
}

func InsertParameters() Params {
	url:= flag.String("url", "", "write your url, write something like: --url http://example.com")
	requests := flag.Int("requests", 0, "number of requests, (--requests > 0), write something like: --requests 1000")
	concurrencies := flag.Int("concurrencies", 0, "number of concurrencies, (--concurrencies > 0) write something like: --concurrencies 1000")

	flag.Parse()

	parameters := Params{
		URL:          *url,
		Requests:     *requests,
		Concurrencies: *concurrencies,
	}

	return parameters
}
