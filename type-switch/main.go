package main

import (
	"fmt"
	"reflect"
)

func main() {
	t1 := "this is a string"
	t2 := &t1
	detectType(t1)
	detectType(t2)
	detectType(123)
	detectType(nil)
	detectType(1.2)
}

func detectType(t1 any) {
	switch v := t1.(type) {
	case string:
		fmt.Printf("Found a string: %v\n", v)
	case *string:
		fmt.Printf("Found a string pointer: %v\n", *v)
	case int:
		fmt.Printf("Found an integer: %v\n", v)
	case nil:
		fmt.Println("Found a nil")
	default:
		fmt.Printf("Found an unknown type: %v\n", reflect.TypeOf(v))
	}
}
