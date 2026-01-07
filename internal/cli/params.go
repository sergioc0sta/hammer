package cli

import (
	"errors"
	"flag"
	"fmt"
	"github.com/sergioc0sta/hammer/internal/utils"
)

type Params struct {
	URL          string
	Requests     int
	Concurrency int
}

var ErrInvalidNumber = errors.New("The value of concurrency is more than requests")
var ErrMissingURL = errors.New("Missing required flag: --url")
var ErrMissingRequests = errors.New("Missing required flag value needs to be more than 0: --requests")
var ErrMissingConcurrency = errors.New("Missing required flag or value needs to be more than 0: --concurrency")

func InsertParameters() (*Params, error) {
	url:= flag.String("url", "", "write your url, write something like: --url http://example.com")
	requests := flag.Int("requests", 0, "number of requests, (--requests > 0), write something like: --requests 1000")
	concurrency := flag.Int("concurrency", 0, "number of concurrency, (--concurrency > 0) write something like: --concurrency 1000")

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Example usage:")
		fmt.Fprintln(flag.CommandLine.Output(), "  docker run hammer --url=http://example.com --requests=1000 --concurrency=10")
	}

	flag.Parse()

	if *url == "" {
		return nil, ErrMissingURL
	}
	if *requests == 0 || *requests < 0 {
		return nil, ErrMissingRequests
	}
	if *concurrency == 0 || *concurrency < 0 {
		return nil, ErrMissingConcurrency
	}
	if *concurrency > *requests {
		return nil, ErrInvalidNumber 
	}
	if err := utils.ValidateURL(*url); err != nil {
		return nil, err
	}

	parameters := Params{
		URL:          *url,
		Requests:     *requests,
		Concurrency: *concurrency,
	}

	return &parameters, nil
}
