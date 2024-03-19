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
	// loop while "ok" is true
	for {
		// receive the value and ok from the channel
		v, ok := <-ch
		// if ok is false break out of the loop
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
