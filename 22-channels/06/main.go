package main

import "fmt"

// [3] the sendData function NOW CONVERTS the channel into a send only channel
// but it is ONLY send only INSIDE the sendData Goroutine
func sendData(sendch chan<- int) {
	// [4] we send 10 into the channel
	sendch <- 10
}

func main() {
	// [1] we make a bidirectional channel
	chnl := make(chan int)
	// [2] run a Goroutine
	go sendData(chnl)
	// we can still receive from the channel because it is STILL bidirectional in the MAIN Goroutine
	fmt.Println(<-chnl)
}
