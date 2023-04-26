package main

import (
	"fmt"
	"sync"
)

type mutexType chan struct{}
type mutex struct {
	s mutexType
}

func NewMutex() mutex {
	return mutex{s: make(mutexType, 1)}
}

func (m mutex) Lock() {
	e := struct{}{}
	m.s <- e
}

func (m mutex) Unlock() {
	<-m.s
}

const N = 1000

func main() {
	m := NewMutex()
	counter := 0
	wg := sync.WaitGroup{}

	wg.Add(N)
	for i := 1; i <= N; i++ {
		go func() {
			m.Lock()
			counter++
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Mutex counter: %d\n", counter)
}
