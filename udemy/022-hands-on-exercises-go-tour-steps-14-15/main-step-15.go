package main

import "fmt"

// Constants

// Constants are declared like variables but with the const keyword
// Cannot be declared using the := syntax
// Can be character, string, boolean, or numeric values

const Pi = 3.14

func main() {

	const World = "Mundo"
	fmt.Println("Hola", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Golang rules?", Truth)
}
