package main

import "fmt"

// In this exercise, x and y are declared outside the function.
// Now, becoming package level variables

var x int = 22
var y string = "Hello Gopher!"

func main() {

	fmt.Println(x)
	fmt.Println(y)
}
