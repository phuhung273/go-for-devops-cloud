package main

import (
	"fmt"
	"time"
)

func main() {
	// goToSleep()

	usingChannel()
}

func usingChannel() {
	fmt.Println("one")
	c := make(chan string)
	go search(c)
	fmt.Println("two")
	result := <-c
	fmt.Printf("Found: %v", result)
}

func goToSleep() {
	fmt.Println("one")
	go fmt.Println("three")
	fmt.Println("two")
	time.Sleep(1 * time.Second)
}

func search(c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Searching...")
		time.Sleep(1 * time.Second)
	}

	c <- "Done searching"
}
