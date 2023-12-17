package main

import "fmt"

// Type inference

// When declaring a variable without a specifying an explicit type
// (either by using the := syntax or var = expression syntax),
// the variable's type will be inferred from the value on the
// right hand side.

func main() {
	v := 42
	w := 3.142
	y := 0.867 + 0.5i

	fmt.Printf("v whose value is %v is of type %T\n", v, v)
	fmt.Printf("w whose value is %v is of type %T\n", w, w)
	fmt.Printf("y whose value is %v is of type %T\n", y, y)
}
