package main

import (
	"fmt"
)

func main() {
	// make a buffered channel (the 2 is the capacity, we can write 2 strings to the channel)
	ch := make(chan string, 2)
	// send a string to the channel
	ch <- "naveen"
	// send another string to the channel
	ch <- "paul"
	// read from the channel and print
	fmt.Println(<-ch)
	// read from the channel and print
	fmt.Println(<-ch)
}
