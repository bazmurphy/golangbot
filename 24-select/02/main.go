package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	// Sleep for 10.5 seconds before sending value to channel
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	ch := make(chan string)
	// Start the `process` function as a goroutine
	go process(ch)

	for {
		// Sleep for 1 second at the start of each iteration
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			// If a value is received from the channel, print it and exit
			fmt.Println("received value: ", v)
			return
		default:
			// If no value is received from the channel, print this message
			fmt.Println("no value received")
		}
	}
}

// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// no value received
// received value:  process successful
