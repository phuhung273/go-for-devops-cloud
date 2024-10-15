package main

import (
	"fmt"
	"reflect"
)

func main() {
	t1 := 123
	fmt.Printf("Result: %v, type: %v\n", plusOne(t1), reflect.TypeOf(plusOne(t1)))
	t2 := 123.23
	fmt.Printf("Result: %v, type: %v\n", plusOne(t2), reflect.TypeOf(plusOne(t2)))
	fmt.Printf("Result: %v, type: %v\n", sum(t1, t1), reflect.TypeOf(sum(t1, t1)))
	fmt.Printf("Result: %v, type: %v\n", sum(t2, t2), reflect.TypeOf(sum(t2, t2)))
	// fmt.Printf("Result: %v, type: %v\n", sum(t1, t2), reflect.TypeOf(sum(t1, t2))) # This wont work
}

func plusOne[V int | int64 | int32 | float32 | float64](t V) V {
	return t + 1
}

// Below won work
func sum[V int | int64 | int32 | float32 | float64](a V, b V) V {
	return a + b
}

// // Below wont work
// func sum[V1 int | int64 | int32 | float32 | float64, V2 int | int64 | int32 | float32 | float64](a V1, b V2) V1 {
// 	return a + b
// }
