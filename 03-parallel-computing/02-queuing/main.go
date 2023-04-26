package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 3
const MESSAGES = 10

func main() {
	var wg sync.WaitGroup

	fmt.Println("Queue of length N:", N)
	queue := make(chan struct{}, N)

	wg.Add(MESSAGES)

	for w := 1; w <= MESSAGES; w++ {
		process(w, queue, &wg)
	}

	wg.Wait()

	close(queue)
	fmt.Println("Processing completed")
}

func process(payload int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()

		fmt.Printf("Start processing of %d\n", payload)
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Completed processing of %d\n", payload)

		<-queue
	}()
}
