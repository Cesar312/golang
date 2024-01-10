package main

import "fmt"

// Ranges

func main() {

	var slice []string
	slice = append(slice, "a", "b", "c", "d")

	for i, v := range slice {
		fmt.Println(i, v)
	}

	// Print the 3rd element of the slice
	fmt.Println("\n", slice[2]) // Added newline before printing element
}
