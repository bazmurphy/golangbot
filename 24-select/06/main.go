package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	// send "from server1" to the channel
	ch <- "from server1"
}

func server2(ch chan string) {
	// send "from server2" to the channel
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	// start the server1 goroutine
	go server1(output1)
	// start the server2 goroutine
	go server2(output2)

	// sleep for 1 second to allow the goroutines to execute
	time.Sleep(1 * time.Second)

	// Random selection
	// When multiple cases in a `select` statement are ready, one of them will be executed at random.
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

// from server1
// OR (random)
// from server2
