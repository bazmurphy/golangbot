package main

import (
	"fmt"
	"sync"
)

// shared variable that will be incremented
var x = 0

func increment(wg *sync.WaitGroup) {
	// increment the shared variable x
	x = x + 1
	// notify the WaitGroup that this goroutine has completed
	wg.Done()
}

func main() {
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		// start a new goroutine to call the increment function
		go increment(&w)
	}

	// wait for all goroutines to complete
	w.Wait()
	fmt.Println("final value of x", x)
}
