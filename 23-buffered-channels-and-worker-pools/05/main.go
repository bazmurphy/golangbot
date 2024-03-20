package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	// The for range loop will read all the values written to the channel
	// and will quit once there are no more values to read since the channel is already closed.
	for n := range ch {
		fmt.Println("Received:", n)
	}
}

// Received: 5
// Received: 6
