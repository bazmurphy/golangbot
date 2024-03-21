package main

func main() {
	// We know that the select statement will block until one of its cases is executed.
	// In this case, the select statement doesn’t have any cases and hence it will block forever resulting in a deadlock.
	// This program will panic
	select {}
}

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [select (no cases)]:
// main.main()
