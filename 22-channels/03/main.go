package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	// [4] print
	fmt.Println("hello go routine is going to sleep")
	// [5] wait 4 seconds
	time.Sleep(4 * time.Second)
	// [6] print
	fmt.Println("hello go routine awake and going to write to done")
	// [7] send a boolean to the channel
	done <- true
}

func main() {
	// [1] make the channel
	done := make(chan bool)
	// [2] print
	fmt.Println("Main going to call hello go goroutine")
	// [3] run the go routine passing it the channel [BLOCK]
	go hello(done)
	// [8] receive the boolean from the channel [UNBLOCK]
	<-done
	// [9] print
	fmt.Println("Main received data")
}

// Main going to call hello go goroutine
// hello go routine is going to sleep
// hello go routine awake and going to write to done
// Main received data
