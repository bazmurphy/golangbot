package main

import (
	"fmt"
)

// takes in a channel as a second argument
func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		// send the digit to the channel
		dchnl <- digit
		number /= 10
	}
	// close the channel
	close(dchnl)
}

// takes in a channel as a second argument
func calcSquares(number int, squareop chan int) {
	sum := 0
	// make a new channel
	dch := make(chan int)
	// run a Goroutine
	go digits(number, dch)
	// receive the digits out of the channel
	for digit := range dch {
		sum += digit * digit
	}
	// send the sum to the channel
	squareop <- sum
}

// takes in a channel as a second argument
func calcCubes(number int, cubeop chan int) {
	sum := 0
	// make a new channel
	dch := make(chan int)
	// run a Goroutine
	go digits(number, dch)
	// recieve the digits out of the channel
	for digit := range dch {
		sum += digit * digit * digit
	}
	// send the sum to the channel
	cubeop <- sum
}

func main() {
	number := 589
	// make two channels
	sqrch := make(chan int)
	cubech := make(chan int)
	// run two Goroutines
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	// receive the values from the two channels
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
