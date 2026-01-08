package main

import (
	"fmt"
	"sync"

	"github.com/sergioc0sta/hammer/internal/cli"
	"github.com/sergioc0sta/hammer/internal/infra/handler"
)

func main() {

	var wg sync.WaitGroup
	prs, err := cli.InsertParameters()

	if err != nil {
		panic(err)
	}

	fmt.Println("URL:", prs.URL)
	fmt.Println("Requests:", prs.Requests)
	fmt.Println("Concurrency:", prs.Concurrency)
	fmt.Println("Server is starting...")

	for i := 0; i < prs.Concurrency; i++ {

		wg.Add(1)
		go func(numberRequests int) {
			defer wg.Done()

			for i := 0; i < numberRequests; i++ {
				handler.ServerRequestHandler(prs.URL)

			}
		}(prs.Requests)
	}

	wg.Wait()
}
