package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for { // listen for the context to be canceled
		select {
		case <-ctx.Done():
			fmt.Println("context is done:", ctx.Err())
			return
		default:
			fmt.Println("working...")          // this will be printed every 500ms until the context is canceled (4 times for 2 seconds sleep in line 28)
			time.Sleep(500 * time.Millisecond) // simulate work taking 500ms
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go doWork(ctx) // start work in a goroutine

	time.Sleep(2 * time.Second) // allow time for work to be done

	fmt.Println("canceling the context...")
	cancel() // cancel the work

	time.Sleep(1 * time.Second) // allow time for work to be canceled
	fmt.Println("done")
}
