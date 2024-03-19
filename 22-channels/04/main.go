package main

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		fmt.Println("calcSquares, sum", digit)
		number /= 10
	}
	// send the sum to the channel
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		fmt.Println("calcCubes, sum", sum)
		number /= 10
	}
	// send the sum to the channel
	cubeop <- sum
}

func main() {
	number := 589
	// make two channels of type int
	sqrch := make(chan int)
	cubech := make(chan int)
	// run two goroutines giving both the number and each their own channel [BLOCK]
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	// receive from the channels into the two respective variables [UNBLOCK]
	squares, cubes := <-sqrch, <-cubech
	// print
	fmt.Println("Final output", squares+cubes)
}
