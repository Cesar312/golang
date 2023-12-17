package main

import "fmt"

// Short variable declarations
// := short assignment statement can be used in place of a
// var declaration with implicit t

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
