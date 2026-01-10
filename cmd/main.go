package main

import (
	"fmt"
	"sync"

	"github.com/sergioc0sta/hammer/internal/cli"
	"github.com/sergioc0sta/hammer/internal/dto"
	"github.com/sergioc0sta/hammer/internal/infra/http"
	"github.com/sergioc0sta/hammer/internal/infra/report"
	"github.com/sergioc0sta/hammer/internal/infra/timer"
	"github.com/sergioc0sta/hammer/internal/infra/worker"
)

func main() {
	var wg sync.WaitGroup
	prs, err := cli.InsertParameters()

	if err != nil {
		panic(err)
	}

	cha := make(chan dto.WorkerResult, prs.Concurrency)
	httpClient := http.NewClient()
	fmt.Println("Server is running...")

	timer := timer.NewTimer()
	wg.Add(prs.Concurrency)
	for i := 0; i < prs.Concurrency; i++ {
		go worker.Worker(prs.Requests, prs.URL, &wg, cha, httpClient)
	}

	go func() {
		wg.Wait()
		close(cha)
	}()

	loadTest := report.NewReport(cha)
	loadTest.Collect()
	loadTest.Print()
	timer.TimeClose()
	fmt.Printf("Total time taken: %s\n", timer.Duration)

}
