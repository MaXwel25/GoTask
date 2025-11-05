package main

import "fmt"

type copyer interface {
	string | int
}

func copyfy[T copyer](slice []T) []T {
	copySlice := make([]T, len(slice))

	copy(copySlice, slice)
	return copySlice
}

func main() {
	stringSlice := []string{"apple", "banana", "cherry"}
	stringCopy := copyfy(stringSlice)

	fmt.Printf("Original string slice: %v\n", stringSlice)
	fmt.Printf("Copied string slice: %v\n", stringCopy)

	intSlice := []int{10, 20, 30}
	intCopy := copyfy(intSlice)

	fmt.Printf("Original int slice: %v\n", intSlice)
	fmt.Printf("Copied int slace: %v\n", intCopy)
}

