package main

import (
	"fmt"
	"sync"
)

const totalJobs = 4
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
	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)

		go func(job int) {
			defer wg.Done()

			fmt.Printf("Worker %d started job %d\n", id, job)

			// Do work and send result
			result := job * 2
			results <- result

			fmt.Printf("Worker %d finished job %d\n", id, job)
		}(j)
	}

	wg.Wait()
}
