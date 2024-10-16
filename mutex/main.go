package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type MyType struct {
	Counter int
	mu      sync.Mutex
}

func main() {
	t := MyType{}

	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(t *MyType) {
			t.mu.Lock()
			fmt.Printf("Input: %d\n", t.Counter)
			t.Counter++
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			fmt.Printf("Output: %d\n", t.Counter)
			c <- true
			t.mu.Unlock()
		}(&t)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}
