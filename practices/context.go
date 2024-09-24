package main

/*
	- Context is a package that provides a way to pass request-scoped values, cancellation signals, and deadlines across API boundaries and between processes.
	- a Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	- Context is often used when a function call needs to be canceled or have a deadline.
	- Context is a part of the standard library, and it is used in many packages.
	- A context.Context is created for each request by the net/http machinery. And is available with the Context() method.
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
