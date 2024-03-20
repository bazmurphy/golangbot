package main

import (
	"fmt"
	"time"
)

// takes a channel as an argument
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		// sends i to the channel
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	// closes the channel
	close(ch)
}

func main() {
	// make a new channel of type int with capacity 2 (buffered channel)
	ch := make(chan int, 2)
	// run a Goroutine
	go write(ch)
	time.Sleep(2 * time.Second)
	// receive the values from the channel
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		// sleep for 2 seconds, this blocks the main Goroutine and any further reading from the channel
		time.Sleep(2 * time.Second)
	}
}
