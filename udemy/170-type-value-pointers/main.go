package main

import "fmt"

func main() {

	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v \t %T\n", &x, &x) // Print the value and type of pointer x

}
