package main

import "fmt"

// Uncomment line 16 before running this file

func main() {

	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	y := &x                        // Assign a pointer variable to the memory address of x
	fmt.Printf("%v \t %T\n", y, y) // Print the value and type of the pointer variable y
	fmt.Println(*y)                // Dereference the pointer variable
	// fmt.Println(*&x)               // Simply x (Note the staticcheck)

	*y = 43        // Assigning a different value
	fmt.Println(x) // Note the change to x
	fmt.Println(*y)

}
