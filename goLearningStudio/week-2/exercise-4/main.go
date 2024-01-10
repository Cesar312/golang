package main

import "fmt"

// Making a copy of slice

func main() {

	var slice []string
	slice = append(slice, "Hello")

	slicecopy := make([]string, len(slice))
	copy(slicecopy, slice) // Copy slice by calling destination first then the source

	fmt.Println("Original slice: ", slice)
	fmt.Println("Copy of slice: ", slicecopy)

}
