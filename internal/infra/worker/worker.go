package worker

import (
	"sync"

	"github.com/sergioc0sta/hammer/internal/dto"
	"github.com/sergioc0sta/hammer/internal/infra/handler"
)

func Worker(numberRequests int, url string, wg *sync.WaitGroup, cha chan dto.WorkerResult) {
	soma := 0
	countByStatus := make(map[int]int)
	defer wg.Done()

	for soma < numberRequests {
		status, err := handler.ServerRequestHandler(url)
		if err != nil {
			countByStatus[0]++
			soma++
			continue
		}
		countByStatus[status]++
		soma++
	}

	counterResults := make([]dto.Result, 0, len(countByStatus))

	for status, count := range countByStatus {
		counterResults = append(counterResults, dto.Result{
			Status: status,
			Count:  count,
		})
	}

	cha <- dto.WorkerResult{Results: counterResults}
}
