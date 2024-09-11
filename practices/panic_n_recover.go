package main

import "fmt"

func panicHere() {
	panic("Panic!")
	// panic is a built-in function that stops the ordinary flow of control and begins panicking.
}

func main() {
	defer func() {
		if r := recover(); r != nil { // If recover is called outside a deferred function or when the goroutine is not panicking, it returns nil.
			fmt.Println("Recovered from panic:", r) // the recover function in Go is specifically designed to catch and handle panics.
		}
	}()

	panicHere()

	fmt.Println("This line will not be executed.")
}
