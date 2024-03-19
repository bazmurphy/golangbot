package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		// send i into the channel
		chnl <- i
	}
	// close the channel
	close(chnl)
}

func main() {
	// make a channel of type int
	ch := make(chan int)
	// run a Goroutine passing it the channel
	go producer(ch)
	// this time use the "range" keyword (instead of v, ok := <-ch)
	// it will automatically read from the channel until it is closed
	for v := range ch {
		fmt.Println("Received ", v)
	}
}
