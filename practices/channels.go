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
	// if we don't specify the buffer size (10), the channel will be unbuffered and block until the receiver is ready to receive the value
	numChan := make(chan []int, 10)
	var numbers []int
	for i := 0; i < 10; i++ {
		go func() {
			numbers = append(numbers, i) // append the value of i to the slice
			fmt.Printf("Sending %d to the channel...\n", i)
		}()
	}

	time.Sleep(1 * time.Second) // wait for the goroutines to finish

	numChan <- numbers // send the slice of integers to the channel
	nums := <-numChan  // receive slice of integers from the channel

	fmt.Println(nums)
}
