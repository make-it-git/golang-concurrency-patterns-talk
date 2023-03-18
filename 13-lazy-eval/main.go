package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func Make(f func() int) LazyInt {
	var v int
	var once sync.Once
	return func() int {
		once.Do(func() {
			v = f()
			f = nil // so that f can now be GC'ed
		})
		return v
	}
}

func main() {
	n := Make(func() int {
		fmt.Println("Doing expensive calculations")
		return 23
	})
	fmt.Println(n())      // Calculates the 23
	fmt.Println(n() + 42) // Reuses the calculated value
}
