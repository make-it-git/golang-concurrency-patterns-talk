package main

import (
	"fmt"
)

const totalJobs = 10
const totalWorkers = 2

func main() {
	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	for w := 1; w <= totalWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}

	close(jobs)

	// Receive results
	for a := 1; a <= totalJobs; a++ {
		<-results
	}

	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		go func(job int) {
			fmt.Printf("Worker %d started job %d\n", id, job)

			// Do work and send result
			result := job * 2
			results <- result

			fmt.Printf("Worker %d finished job %d\n", id, job)
		}(j)
	}
}
