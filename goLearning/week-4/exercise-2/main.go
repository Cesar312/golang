package main

// Concurrency with channels

import (
	"fmt"
	// "time"
)

func response(responses []string, ch chan string) {
	for i, item := range responses {
		fmt.Println("Print from insider the function", i, item)
		ch <- item
	}
	close(ch)
}

func main() {
	greetings1 := []string{"Hi", "How are you?", "Great!", "OK, see you later!"}
	ch1 := make(chan string)
	go response(greetings1, ch1)
	for c := range ch1 {
		fmt.Println("Print from inside channel-1", c)
	}

	greetings2 := []string{"Hola", "How is it going?", "Wonderful!", "Goodbye!"}
	ch2 := make(chan string)
	go response(greetings2, ch2)
	for c := range ch2 {
		fmt.Println("Print from inside channel-2", c)
	}
}
