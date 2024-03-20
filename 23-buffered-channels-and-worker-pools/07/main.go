package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	// The counter is decremented by the call to wg.Done() in the process Goroutine.
	// Once all the 3 spawned Goroutines finish their execution,
	// that is once wg.Done() has been called three times,
	// the counter will become zero,
	// and the main Goroutine will be unblocked.
	wg.Done()
}

func main() {
	no := 3
	// WaitGroup is a struct type and we are creating a zero value variable of type WaitGroup
	// The way WaitGroup works is by using a counter.
	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		// When we call wg.Add() on the WaitGroup and pass it an int, the WaitGroup's counter is incremented by the value passed to Add().
		wg.Add(1)
		// It is important to pass the pointer of wg (&wg)
		// If the pointer is not passed, then each Goroutine will have its own copy of the WaitGroup
		// and main will not be notified when they finish executing.
		go process(i, &wg)
	}

	// The Wait() method blocks the Goroutine in which it's called until the counter becomes zero.
	wg.Wait()

	fmt.Println("All go routines finished executing")
}

// started Goroutine  2
// started Goroutine  1
// started Goroutine  0
// Goroutine 0 ended
// Goroutine 2 ended
// Goroutine 1 ended
// All go routines finished executing
