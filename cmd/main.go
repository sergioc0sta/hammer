package main

import (
	"fmt"
	"sync"

	"github.com/sergioc0sta/hammer/internal/cli"
	"github.com/sergioc0sta/hammer/internal/dto"
	"github.com/sergioc0sta/hammer/internal/infra/worker"
)

func main() {
	totalRequests := 0
	var wg sync.WaitGroup
	prs, err := cli.InsertParameters()

	if err != nil {
		panic(err)
	}

	cha := make(chan dto.WorkerResult, prs.Concurrency)
	fmt.Println("URL:", prs.URL)
	fmt.Println("Requests:", prs.Requests)
	fmt.Println("Concurrency:", prs.Concurrency)
	fmt.Println("Server is starting...")

	wg.Add(prs.Concurrency)
	for i := 0; i < prs.Concurrency; i++ {
		go worker.Worker(prs.Requests, prs.URL, &wg, cha)
	}

	go func() {
		wg.Wait()
		close(cha)
	}()

	for v := range cha {
		for _, result := range v.Results {
			fmt.Printf("Status: %d - Count: %d\n", result.Status, result.Count)
			totalRequests += result.Count
		}
	}

	fmt.Println("Total requests:", totalRequests)

}
