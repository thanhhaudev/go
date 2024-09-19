package main

/*
Notes:
- In Go, channels are used to communicate between goroutines.
- A channel is a communication pipe that allows one goroutine to send data to another goroutine.
- By default, channels are unbuffered, meaning that they block the sender until the receiver is ready to receive the data.
- Buffered channels have a fixed capacity that allows them to store a certain number of values without blocking the sender.
*/

import (
	"fmt"
	"time"
)

func main() {
	// this is an unbuffered channel
	messages := make(chan string)

	go func() { // anonymous function, "go" keyword creates a new goroutine
		messages <- "ping" // send a value to the channel
	}()

	msg := <-messages // receive a value from the channel

	// Messages can be sent and received in any order.
	// If the channel is empty, the receiver will block until a value is sent.
	fmt.Println(msg)

	// this is buffered channel
	// if we don't specify the buffer size (2), the channel will be unbuffered and block until the receiver is ready to receive the value
	numChan := make(chan []int, 2)

	go func() {
		numChan <- []int{1, 2, 3, 4, 5} // send a slice of integers to the channel
		fmt.Println("Sent first slice of integers to the channel.")
	}()

	go func() {
		numChan <- []int{6, 7, 8, 9, 10} // send another slice of integers to the channel
		fmt.Println("Sent second slice of integers to the channel.")
	}()

	time.Sleep(1 * time.Second) // wait for the goroutines to finish

	close(numChan) // close the channel after sending the values
	for num := range numChan {
		fmt.Printf("Received slice of integers: %v\n", num)
	}
}
