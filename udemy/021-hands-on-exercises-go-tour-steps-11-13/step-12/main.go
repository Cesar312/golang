package main

import "fmt"

// Zero value

// Variables declared withour an explicit intial value are given
// their zero value

func main() {

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
