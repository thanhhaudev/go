package main

import (
	"context"
	"fmt"
	"time"
)

/*
context.WithTimeout is a shorthand for context.WithDeadline(parent, time.Now().Add(timeout)).

Difference between context.WithTimeout and context.WithDeadline:
 - context.WithTimeout is used when you want to cancel the context when the timeout is reached.
 - context.WithDeadline is used when you want to cancel the context when the deadline is reached.
 - context.WithTimeout is used with a time duration, context.WithDeadline is used with a specific time.
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // create a context with a timeout of 2 seconds
	defer cancel()

	// create a goroutine to simulate a task that takes 3 seconds to complete
	go func() {
		fmt.Println("Start task...")

		// simulate a task that takes 3 seconds to complete
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Task completed.")
		case <-ctx.Done():
			fmt.Println("Task canceled:", ctx.Err())
		}
	}()

	// wait for the task to complete or the context to be canceled
	time.Sleep(4 * time.Second) // wait for 4 seconds to ensure the task is completed or canceled, if we don't wait, the main goroutine will exit before the task is completed or canceled
	fmt.Println("Done.")
}
