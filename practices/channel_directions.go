package main

import (
	"fmt"
	"time"
)

/*
Channel Directions in Go are used to specify the direction of data flow in a channel.
There are three types of channel directions in Go:
1. Send-Only Channel: chan<- T
2. Receive-Only Channel: <-chan T
3. Bidirectional Channel (Send-Receive): chan T

Control how a function can use a channel to send or receive data helps to prevent misuse of the channel and increase the safety of the program.
*/

// Receiving data from the channel
func receiveOnlyChannel(ch <-chan int) {
	fmt.Println(<-ch)
}

// Sending data to the channel
func sendOnlyChannel(ch chan<- int, value int) {
	ch <- value
}

// Send and receive data from the channel, can understand as a both send-only and receive-only channel
func bidirectionalChannel(ch chan int, value int) {
	ch <- value
	fmt.Println(<-ch)
}

func main() {
	// "Channel Directions" mean we control how a function can use a channel to send or receive data.
	ch := make(chan int)

	// Send data to the channel
	go sendOnlyChannel(ch, 1)

	// Receive data from the channel
	go receiveOnlyChannel(ch)

	// Send and receive data from the channel
	go bidirectionalChannel(ch, 2)

	time.Sleep(1 * time.Second) // because we don't receive data from the channel, then it does not wait for all goroutine to finish.
	// if we receive data from the channel, then it will wait for all goroutine to finish. It's because the main goroutine is blocked while waiting to receive data from the channel.

	fmt.Println("Finished")
}
