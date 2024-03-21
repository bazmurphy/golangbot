package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	// Sleeps for 6 seconds
	time.Sleep(6 * time.Second)

	// Writes the text "from server1" to the channel `ch`
	ch <- "from server1"
}

func server2(ch chan string) {
	// Sleeps for 3 seconds
	time.Sleep(3 * time.Second)

	// Writes the text "from server2" to the channel `ch`
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	// Calls the server1 function in a new Goroutine
	go server1(output1)

	// Calls the server2 function in a new Goroutine
	go server2(output2)

	select {
	// The `select` statement blocks until one of its cases is ready
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
