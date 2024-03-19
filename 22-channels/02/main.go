package main

import (
	"fmt"
)

// hello takes a channel as an argument
func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	// send true to the done channel
	done <- true
}

func main() {
	// create a done channel (of type boolean)
	done := make(chan bool)
	// pass the done channel to the hello go routine
	// this line of code is blocking
	// until some Goroutine writes data the `done` channel the control will not move to the next line of code
	go hello(done)
	// receives data from the done channel but does not use or store that data in any variable. This is perfectly legal.
	<-done
	// this print line is blocked until main receives data from the done channel
	fmt.Println("main function")
}

// Hello world goroutine
// main function
