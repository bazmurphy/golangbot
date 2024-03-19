package main

func main() {
	// create a channel
	ch := make(chan int)
	// send 5 to the channel
	ch <- 5
	// but no other Goroutine is recieving data from the channel
	// so the program panics:
	// "fatal error: all goroutines are asleep - deadlock!"
}
