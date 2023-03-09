package main

import (
	"fmt"
	"sync"
)

type semaphore chan struct{}
type mutex struct {
	s semaphore
}

func NewSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func NewMutex() mutex {
	return mutex{s: make(semaphore, 1)}
}

func (s semaphore) Acquire(n int) {
	e := struct{}{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func (m mutex) Lock() {
	m.s.Acquire(1)
}

func (m mutex) Unlock() {
	m.s.Release(1)
}

const N = 1000

func main() {
	m := NewMutex()
	mutexCounter := 0
	wg := sync.WaitGroup{}

	wg.Add(N)
	for i := 1; i <= N; i++ {
		go func() {
			m.Lock()
			mutexCounter++
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Mutex counter: %d\n", mutexCounter)
}
