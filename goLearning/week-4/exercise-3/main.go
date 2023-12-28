package main

// Concurrency with sync.WaitGroup

import (
	"fmt"
	"sync"
)

func respond(greeting string, ch chan string) {
	ch <- greeting
}

func main() {

	greeting := []string{"Hi", "How are you?", "Great!", "OK, see you later!"}
	var wg sync.WaitGroup

	ch := make(chan string)
	wg.Add(2)
	go respond(greeting[1], ch)
	go respond(greeting[2], ch)
	wg.Done()
	v := <-ch
	fmt.Println(v)
	wg.Done()
	v = <-ch
	fmt.Println(v)
	wg.Wait()
	fmt.Println("Program finished")

}
