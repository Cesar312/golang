package main

import "fmt"

// NOTE: similar to step-4 file but simplified version
// 		 by sharing the same type
// Function to add two arguments

func add(x, y int) int {
	return x + y
}

// Print added argument in the console

func main() {
	fmt.Println(add(5, 32))
}
