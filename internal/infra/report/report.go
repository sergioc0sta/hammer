package report

import (
	"fmt"

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

func (r *Report) Print() {
	fmt.Println("Total requests:", r.TotalRequests)
	for status, count := range r.CountByStatus {
		fmt.Printf("Status: %d - Count: %d\n", status, count)
	}
}
