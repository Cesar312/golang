package main

import "fmt"

func main() {

	var textOut string
	var myName = "Ivan"

	textOut = fmt.Sprintf("Hi, %v. Welcome!", myName)
	fmt.Println(textOut)

}
