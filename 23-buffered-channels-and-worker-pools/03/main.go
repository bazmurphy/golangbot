package main

import (
	"fmt"
)

func main() {
	// we make a buffered channel of capacity 2
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	// so when the control reaches the third write
	// the write is blocked since the channel has exceeded its capacity
	ch <- "steve"
	// some Goroutine needs to read from the channel to unblock it so the write can proceed
	// so there is a deadlock and the program panics
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan send]:
