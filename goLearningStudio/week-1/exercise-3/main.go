package main

import "fmt"

func main() {

	x := 22 // Short variable declaration
	fmt.Println(x)

	// Reusing x variable BUT with DIFFERENT operator
	x = 44 // Can NOT reuse short variable declaration on the same variable
	fmt.Println(x)

}
