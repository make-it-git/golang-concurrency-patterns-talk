package main

import (
	"fmt"
	"sync"
	"time"
)

type payload struct {
	name  string
	value int
}

func producer(name string, done <-chan struct{}, wg *sync.WaitGroup) <-chan payload {
	ch := make(chan payload)
	var i = 1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(ch)
				fmt.Println(name, "completed")
				return
			case ch <- payload{
				name:  name,
				value: i,
			}:
				fmt.Println(name, "produced", i)
				i++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	return ch
}

func consumer(name string, channels []<-chan payload, done <-chan struct{}, wg *sync.WaitGroup) {
	for i, ch := range channels {
		i := i + 1
		ch := ch
		go func() {
			defer wg.Done()
			fmt.Println("started consumer", name, i)
			for {
				select {
				case <-done:
					fmt.Println("consumer", name, i, "completed")
					return
				case v := <-ch:
					fmt.Println("Consumer", name, i, "got value", v.value, "from", v.name)
				}
			}
		}()
	}
}

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(3)
	producers := make([]<-chan payload, 0, 3)
	producers = append(producers, producer("Alice", done, &wg))
	producers = append(producers, producer("Jack", done, &wg))
	producers = append(producers, producer("Bob", done, &wg))

	wg.Add(3)
	consumer("C1", producers, done, &wg)

	wg.Add(3)
	consumer("C2", producers, done, &wg)

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}
