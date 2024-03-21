package main

import (
	"fmt"
	"sync"
)

// shared variable that will be incremented
var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	// acquire the mutex lock to ensure exclusive access to the critical section
	m.Lock()
	// critical section: increment the shared variable x
	x = x + 1
	// release the mutex lock
	m.Unlock()
	// notify the WaitGroup that this goroutine has completed
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	// create a mutex to control access to the critical section
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		w.Add(1)
		// start a new goroutine to call the increment function, passing the mutex
		go increment(&w, &m)
	}

	// wait for all goroutines to complete
	w.Wait()
	fmt.Println("final value of x", x)
}

// final value of x 1000
