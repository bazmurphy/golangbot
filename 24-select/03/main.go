package main

func main() {
	ch := make(chan string)
	select {
	// If a default case is present, this deadlock will not happen
	// since the default case will be executed when no other case is ready.
	case <-ch:
	}
}

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
