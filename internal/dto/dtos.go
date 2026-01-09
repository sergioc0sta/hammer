package dto

type Result struct {
	Status int
	Count  int
}

type WorkerResult struct {
	Results []Result
}
