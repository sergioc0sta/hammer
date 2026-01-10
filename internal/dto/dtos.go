package dto

type Result struct {
	Status int
	Count  int
}

type WorkerResult struct {
	Results []Result
}

type CountByStatus map[int]int
