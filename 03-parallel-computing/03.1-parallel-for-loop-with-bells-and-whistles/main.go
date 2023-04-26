package main

import (
	"fmt"
	"math/rand"
	"time"
)

const DATA_SIZE = 4

func calculate(val int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return val * 2
}

func main() {
	data := make([]int, 0, DATA_SIZE)
	for i := 0; i < DATA_SIZE; i++ {
		data = append(data, i+10)
	}
	results := make([]int, DATA_SIZE)
	semaphore := make(chan int, DATA_SIZE)

	fmt.Printf("Before: %v\n", data)
	start := time.Now()

	for i, xi := range data {
		go func(i int, xi int) {
			results[i] = calculate(xi)
			semaphore <- results[i]
		}(i, xi)
	}
	for i := 0; i < DATA_SIZE; i++ {
		fmt.Printf(" one calculation completed: %d\n", <-semaphore)
	}

	fmt.Printf(" After: %v\n", results)
	fmt.Printf(" Elapsed: %s\n", time.Since(start))
}
