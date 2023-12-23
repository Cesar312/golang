package main

import "fmt"

// Array of strings

func main() {

	var myArray [3]string

	myArray[0] = "Hello"
	myArray[1] = "Gopher"
	myArray[2] = "Hello again, Gopher!"

	fmt.Println(myArray[0]) // Print 1st string in the array
	fmt.Println(myArray)    // Print the entire array
}
