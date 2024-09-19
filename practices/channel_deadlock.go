package main

import "fmt"

/*
Deadlock is a situation that happens when a goroutine sends or receives data not in a proper way to the channel.
Deadlock occurs when goroutines indefinitely wait for each other to finish and release the resources.
*/
func main() { // although the main is a goroutine, it is not a separate goroutine, so the deadlock occurs here
	//unbufferedChannelDeadlock()
	//bufferedChannelDeadlock()

	unbufferedChannel()
	bufferedChannel()

	// why in this file, we don't need to use time.Sleep() to wait for the goroutines to finish?
	// because the main goroutine is blocked while waiting to receive data from the channel
	// the main goroutine will only proceed once it has received the data, which ensures that the other goroutines have had a chance to send their data.
}

// this function try to send and receive with the same routine (main), so the deadlock occurs
func unbufferedChannelDeadlock() {
	// Create an unbuffered channel
	ch := make(chan int)
	// unbuffered channel requires a sender and a receiver to be ready at the same time, otherwise, deadlock occurs

	ch <- 1 // deadlock occurs here because the receiver is not ready to receive the value, we need to create a goroutine to send the value to the channel, e.g., go func() { ch <- 1 }()
	// unbuffered channel deadlock occurs when:
	// - there is no receiver to receive the value from the channel
	// ch := make(chan int)
	// ch <- 1
	// - there is no sender to send the value to the channel
	// ch := make(chan int)
	// <-ch

	fmt.Println(<-ch) // Receive the value from the channel
}

// this function try to send and receive with the same routine (main), so the deadlock occurs
func bufferedChannelDeadlock() {
	// Create a buffered channel
	ch := make(chan int, 2)
	// buffered channel does not require a sender and a receiver to be ready at the same time, so we can send many values to the channel without receiving them immediately

	ch <- 1 // sent
	ch <- 2 // sent
	ch <- 3 // deadlock occurs here because the buffer is full, we need to receive the value from the channel to free the buffer, e.g., fmt.Println(<-ch)
	// buffered channel deadlock occurs when:
	// - the buffer is full and the sender tries to send more values to the channel
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// ch <- 3
	// - the buffer is empty and the receiver tries to receive the value from the channel
	// ch := make(chan int, 2)
	// fmt.Println(<-ch)

	fmt.Println(<-ch) // Receive the value from the channel
}

func unbufferedChannel() {
	// Create an unbuffered channel
	ch := make(chan int)

	go func() {
		ch <- 0 // Send a value to the channel
	}()

	fmt.Println(<-ch) // Receive the value from the channel
}

func bufferedChannel() {
	// Create a buffered channel
	ch := make(chan int, 2)

	go func() {
		ch <- 1 // Send a value to the channel
		ch <- 2 // Send another value to the channel
	}()

	fmt.Println(<-ch) // Receive the value from the channel
	fmt.Println(<-ch) // Receive another value from the channel
}
