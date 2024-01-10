package main

import "fmt"

// Slices
// Several examples in this exercise
// Uncomment sections to view outputs

func main() {

	var s []int

	// Section - 1

	// fmt.Println(s, len(s)) // Print slice and its length

	// Section - 2

	s = append(s, 1) // Append a single element to slice
	fmt.Println(s)   // Print slice

	// Section - 3

	// s = append(s, 2, 3, 4, 5, 6) // Append multiple elements to slice
	// fmt.Println(s)               // Print slice

}
