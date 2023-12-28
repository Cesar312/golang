package main

// Concurrency

import (
	"fmt"
	"time"
)

func response(responses []string) {
	for i, item := range responses {
		fmt.Println(i, item)
	}
}

func main() {
	greetings1 := []string{"Hi", "How are you?", "Great!", "OK, see you later!"}
	greetings2 := []string{"Hola", "How is it going?", "Wonderful!", "Goodbye!"}

	go response(greetings1)
	go response(greetings2)

	time.Sleep(3 * time.Second)
	fmt.Println("Program ended")
}
