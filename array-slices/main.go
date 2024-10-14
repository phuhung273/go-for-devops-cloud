package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1)
	fmt.Printf("Len: %d, Cap: %d\n", len(arr1), cap(arr1))
	arr2 := arr1[1:4]
	fmt.Println(arr2)
	fmt.Printf("Len: %d, Cap: %d\n", len(arr2), cap(arr2))

	// See the underlying array of slice
	arr2 = arr2[0 : len(arr2)+2]
	fmt.Println(arr2)
	fmt.Printf("Len: %d, Cap: %d\n", len(arr2), cap(arr2))

	// Modifying slice also change underlying array
	for i := range arr2 {
		arr2[i] += 1
	}
	fmt.Println(arr1)

	// Slice increase capacity if reached
	arr3 := []int{1, 2, 3}
	fmt.Println(arr3)
	fmt.Printf("Len: %d, Cap: %d\n", len(arr3), cap(arr3))
	arr3 = append(arr3, 4)
	fmt.Println(arr3)
	fmt.Printf("Len: %d, Cap: %d\n", len(arr3), cap(arr3))
	arr3 = append(arr3, 5)
	fmt.Println(arr3)
	fmt.Printf("Len: %d, Cap: %d\n", len(arr3), cap(arr3))
}
