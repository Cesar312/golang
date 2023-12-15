package main

import (
	"fmt"
	"math/rand"
)

// Print random number between 0-9
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
