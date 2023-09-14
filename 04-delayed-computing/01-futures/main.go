package main

import (
	"fmt"
	"time"
)

type data struct {
	Body  string
	Error error
}

func doGet(url string) (string, error) {
	time.Sleep(time.Millisecond * 200)
	return fmt.Sprintf("Response of %s", url), nil
}

func future(url string) <-chan data {
	c := make(chan data, 1)

	go func() {
		body, err := doGet(url)

		c <- data{Body: body, Error: err}
	}()

	return c
}

func main() {
	future1 := future("https://example1.com")
	future2 := future("https://example2.com")

	fmt.Println("Requests started")

	response1 := <-future1
	response2 := <-future2

	fmt.Printf("Response 1: %v\n", response1)
	fmt.Printf("Response 2: %v\n", response2)
}
