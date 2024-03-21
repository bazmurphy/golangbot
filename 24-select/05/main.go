package main

import "fmt"

func main() {
	// `ch` is `nil`
	var ch chan string
	select {
	// and we are trying to read from `ch` in the select
	case v := <-ch:
		fmt.Println("received value", v)
	// the default case will be executed even if the select has only `nil` channels.
	default:
		// If the `default` case was not present, the `select` would have blocked forever and caused a deadlock.
		// Since we have a default case inside the select, it will be executed and the program will print
		fmt.Println("default case executed")
	}
}

// default case executed
