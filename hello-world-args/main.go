package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: ./hello_world <name>")
		os.Exit(1)
	}
	fmt.Printf("Hello %v", args[1])
}
