package main

/*
 - ContextWithValue helps to pass request-scoped values across API boundaries and between processes. So no need to create global variables to pass values.
 - Best practices:
 + Only use context values for request-scoped data that follows the request through the call chain.
 + Do not use context values for passing optional parameters to functions.
 + Do not use context values for passing data that is not request-scoped like configuration data, constants, or variables that are not specific to a request
 + Do not use context values for passing a struct with multiple fields. Instead, pass each field individually.
 + Use a custom type for the key to prevent conflicts with other packages that may use the same key type.
*/
import (
	"context"
	"fmt"
	"time"
)

var ctxLanguageKey struct{} // this is best practice to use an empty struct as the key type, because it has zero size and does not allocate memory. It prevents conflicts with other packages that may use the same key type.

func process(ctx context.Context) {
	language, ok := ctx.Value(ctxLanguageKey).(string)
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
	ctx := context.Background()                        // original context
	ctx = context.WithValue(ctx, ctxLanguageKey, "Go") // add a key-value pair to the context

	go process(ctx) // this means that the process function will be executed in a goroutine, and it will have the language key-value pair

	time.Sleep(3 * time.Second) // wait for work to be done
	fmt.Println("done")
}
