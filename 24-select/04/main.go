package main

import "fmt"

func main() {
	ch := make(chan string)
	select {
	case <-ch:
	// If a default case is present, this deadlock will not happen since the default case will be executed when no other case is ready
	default:
		fmt.Println("default case executed")
	}
}

// default case executed
