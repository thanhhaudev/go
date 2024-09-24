package main

/*
 - ContextWithValue helps to pass request-scoped values across API boundaries and between processes. So no need to create global variables to pass values.
*/
import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	language, ok := ctx.Value("language").(string)
	if !ok {
		fmt.Println("language key not found in the context")
		return
	}

	fmt.Println("language is", language)

	// simulate work
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("work is done")
	case <-ctx.Done():
		fmt.Println("work is canceled:", ctx.Err())
	}
}

func main() {
	ctx := context.Background()                    // original context
	ctx = context.WithValue(ctx, "language", "Go") // add a key-value pair to the context

	go process(ctx) // this means that the process function will be executed in a goroutine, and it will have the language key-value pair

	time.Sleep(3 * time.Second) // wait for work to be done
	fmt.Println("done")
}
