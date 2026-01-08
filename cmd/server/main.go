package main

import (
	"fmt"
	"sync"

	"github.com/sergioc0sta/hammer/internal/cli"
	"github.com/sergioc0sta/hammer/internal/infra/handler"
)

func Worker(numberRequests int, url string, wg *sync.WaitGroup, cha chan int) {
	soma := 0
	defer wg.Done()

	for soma < numberRequests {
		handler.ServerRequestHandler(url)
		soma++
	}
	cha <- soma
}

func main() {
	cha := make(chan int, 1)
	totalRequests := 0
	var wg sync.WaitGroup
	prs, err := cli.InsertParameters()

	if err != nil {
		panic(err)
	}

	fmt.Println("URL:", prs.URL)
	fmt.Println("Requests:", prs.Requests)
	fmt.Println("Concurrency:", prs.Concurrency)
	fmt.Println("Server is starting...")

	wg.Add(prs.Concurrency)
	for i := 0; i < prs.Concurrency; i++ {
		go Worker(prs.Requests, prs.URL, &wg, cha)
	}

	go func() {
		wg.Wait()
		close(cha)
	}()

	for v := range cha {
		totalRequests += v
	}

	fmt.Println("Total requests:", totalRequests)

}
