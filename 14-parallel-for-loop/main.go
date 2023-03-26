package main

import (
	"fmt"
	"time"
)

type empty struct{}

const N = 4

func calculate(val int) int {
	time.Sleep(time.Millisecond * 500)
	return val * 2
}

func main() {
	data := make([]int, 0, N)
	for i := 0; i < N; i++ {
		data = append(data, i+10)
	}
	results := make([]int, N)
	semaphore := make(chan empty, N)

	fmt.Printf("Before: %v\n", data)
	start := time.Now()

	for i, xi := range data {
		go func(i int, xi int) {
			results[i] = calculate(xi)
			semaphore <- empty{}
		}(i, xi)
	}
	for i := 0; i < N; i++ {
		<-semaphore
	}

	fmt.Printf(" After: %v\n", results)
	fmt.Printf(" Elapsed: %s\n", time.Since(start))
}
