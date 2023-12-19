package main

// The imported pacakges "fmt" and "furtherexplored"
// are in the 'File Block' scope of this file

import (
	"fmt"
	"golang/udemy/032-understanding-scope/furtherexplored"
)

// x is in the 'Package Block' scope

var x = 42

func main() {

	// x is accessible here
	fmt.Println(x)

	// x is accessible also here
	printMe()

	// y is undefined at this point
	// so inaccesible and the following statement will not work
	// fmt.Println(y)

	// y is in the 'Block' scope
	y := 43
	fmt.Println(y)

	p1 := person{"Procopio"}
	p1.sayHello()

	// variable 'Shadowing'
	x := 32
	fmt.Println(x)

	// package block x is still the same
	printMe()

	furtherexplored.Fascinating()

	// also good to know
	if z := 100; z > 90 {
		fmt.Println(z)
	}

	// at this point, you CANNOT access z here
	// fmt.Println(z)

	/*
		take a look at the 'Comma OK idiom'
		https://go.dev/doc/effective_go#maps
	*/

}

func printMe() {

	// x is accessible here
	fmt.Println(x)
}

// Type person is in 'Package Block' scope
type person struct {
	first string
}

// The method sayHello
// which is attached to VALUES of TYPE person
// is in 'Package Block' scope
func (p person) sayHello() {
	fmt.Println("Hello, my name is ", p.first)
}
