package main

/*
	- The `context` package provides a way to pass request-scoped values, cancellation signals, and deadlines across API boundaries and between processes.
	- It is often used when a function call needs to be canceled or have a deadline.
	- `context` is part of the standard library and is used in many packages.
	- A `context.Context` is created for each request by the `net/http` machinery and is available via the `Context()` method on the `http.Request` object.
*/
import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second): // simulate long operation
		fmt.Fprintf(w, "hello\n") // write response after 10 seconds
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
API boundaries refer to the boundaries between different parts of a program or between different programs. For example, when a client makes a request to a server, the server is a separate program from the client, and the boundary between them is an API boundary.
*/
