package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
 - This simulates a scenario where a request is made to an external service.
 - The external service takes 5 seconds to respond.
 - The request handler has a deadline of 3 seconds.
 - The external service will be canceled after 3 seconds.
 - The request handler will return an error to the client.
 - This is a contrived example to demonstrate how to use context.WithDeadline in an HTTP handler. Help to prevent the server from waiting for a long time and improve the performance of the server.
*/

func main() {
	http.HandleFunc("/process", processDeadline)
	http.ListenAndServe(":8080", nil)
}

func processDeadline(w http.ResponseWriter, rq *http.Request) {
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(rq.Context(), deadline)
	defer cancel()

	err := externalService(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Request processed")
}

func externalService(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second): // if we change this to 2 seconds, the external service will be done before the deadline then the request will be processed successfully
		// simulate work
		fmt.Println("external service done")
		return nil
	case <-ctx.Done():
		fmt.Println("external service canceled:", ctx.Err())
		return ctx.Err()
	}
}
