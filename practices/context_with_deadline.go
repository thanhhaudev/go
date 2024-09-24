package main

/*
 - Use context.WithDeadline when you want to cancel the context when the deadline is reached. Like a task that should be done in a certain time.
 - This ensures that the resources tied to the context are released when the deadline is reached.
*/
import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	go func() {
		fmt.Println("start working...")
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("work goroutine done")
		case <-ctx.Done(): // this will be triggered when the deadline is reached
			fmt.Println("work goroutine canceled:", ctx.Err())
		}
	}()

	time.Sleep(4 * time.Second)
	fmt.Println("Main goroutine done")
}
