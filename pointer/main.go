package main

import "fmt"

func main() {
	a := "change me"
	fmt.Printf("Original: %s\n", a)
	changeString(a)
	fmt.Printf("Without Pointer: %s\n", a)
	changeStringPointer(&a)
	fmt.Printf("Pointer: %s\n", a)

	b := []string{"change me"}
	fmt.Printf("Original: %v\n", b)
	changeSliceElement(b)
	fmt.Printf("Change Slice element: %v\n", b)
	appendSlice(b)
	fmt.Printf("Append Slice: %v\n", b)
	appendSlicePointer(&b)
	fmt.Printf("Append Slice Pointer: %v\n", b)
}

func changeString(input string) {
	input = "changed"
}

func changeStringPointer(input *string) {
	*input = "changed"
}

func changeSliceElement(input []string) {
	input[0] = "changed"
}

func appendSlice(input []string) {
	input = append(input, "another string")
}

func appendSlicePointer(input *[]string) {
	*input = append(*input, "another string")
}
