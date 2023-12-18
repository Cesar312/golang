package main

import "fmt"

// Print verbs to show values and type

func main() {
	a, b, c := "Real Madrid", 20, 41.02

	fmt.Printf("%v is type %T\n", a, a)
	fmt.Printf("%v is type %T\n", b, b)
	fmt.Printf("%v is type %T\n", c, c)
}
