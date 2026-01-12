package report

import (
	"fmt"
	"sort"
	"time"

	"github.com/sergioc0sta/hammer/internal/dto"
)

type Report struct {
	TotalRequests int
	CountByStatus dto.CountByStatus
	cha           <-chan dto.WorkerResult
}

func NewReport(cha <-chan dto.WorkerResult) *Report {
	return &Report{
		TotalRequests: 0,
		CountByStatus: make(dto.CountByStatus),
		cha:           cha,
	}
}

func (r *Report) Collect() {
	for v := range r.cha {
		for _, result := range v.Results {
			r.CountByStatus[result.Status] += result.Count
			r.TotalRequests += result.Count
		}
	}
}

func (r *Report) Print(duration time.Duration) {
	fmt.Println("=== Hammer Report ===")
	fmt.Printf("Total time: %s\n", duration)
	fmt.Printf("Total requests: %d\n", r.TotalRequests)

	success := r.CountByStatus[200]
	fmt.Printf("Successful requests (200): %d\n\n", success)

	fmt.Println("Status code distribution:")

	keys := make([]int, 0, len(r.CountByStatus))
	for status := range r.CountByStatus {
		keys = append(keys, status)
	}
	sort.Ints(keys)

	for _, status := range keys {
		count := r.CountByStatus[status]
		percent := 0.0
		if r.TotalRequests > 0 {
			percent = (float64(count) / float64(r.TotalRequests)) * 100
		}
		fmt.Printf("%d: %d (%.1f%%)\n", status, count, percent)
	}
}
