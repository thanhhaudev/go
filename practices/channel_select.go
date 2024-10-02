package main

import (
	"fmt"
	"time"
)

/**
Select statement in Go is used to wait on multiple channel operations.
It like a switch statement, but each case in a select statement is a communication operation on a channel.
Select statement blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

Use the select statement when:
- Receive data from multiple channels
- Send data to multiple channels
- Handle timeouts for channel operations
*/

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)
	now := time.Now()
	go func() {
		time.Sleep(1 * time.Second)

		select { // send data to chan1 when it's ready
		case chan1 <- 1:
			println("Sent to chan1")
		default:
			println("No value sent")
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)

		select {
		case chan2 <- 2:
			println("Sent to chan2")
		default:
			println("No value sent")
		}
	}()

	go func() {
		/*
			Here's a step-by-step explanation:
			1. The goroutine sleeps for 3 seconds.
			2. After waking up, it sends the value 3 to chan3.
			3. The select statement tries to send 99 to chan3, but it blocks because chan3 is already full.
			4. Meanwhile, the main function is waiting to receive values from chan1, chan2, and chan3.
			5. The main function receives the value 3 from chan3.
		*/

		time.Sleep(3 * time.Second) // longer time
		chan3 <- 3                  // send data to chan3 immediately, so it will be received first by the main goroutine

		select {
		case chan3 <- 99:
			println("Sent to chan3")
		default: // this will be executed because chan3 is full
			println("chan3 is full, cannot send 99 to chan3")
		}
	}()

	fmt.Println("Waiting for goroutines to finish...")

	var i int
	// receive from multiple channels
	// using a loop to prevent the main goroutine from exiting before receiving all values from the channels,
	// if we don't use the loop, the main goroutine will exit after receiving the first value from the channel
	for t := 0; t < 3; t++ {
		select { // await multiple channels, printing each value as it arrives
		case i = <-chan1:
			println("Received from chan1:", i)
		case i = <-chan2:
			println("Received from chan2:", i)
		case i = <-chan3:
			println("Received from chan3:", i)
		}
	}

	fmt.Println("Done after", time.Since(now))

	/*
		Output:
			Waiting for goroutines to finish...
			Sent to chan1
			Received from chan1: 1
			Sent to chan2
			Received from chan2: 2
			chan3 is full, cannot send 99 to chan3
			Received from chan3: 3
			Done after 3.000236187s

		*Time taken: ~3 seconds because the main goroutine waits for the longest goroutine to finish
	*/
}
