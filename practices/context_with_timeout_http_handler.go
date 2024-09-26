package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/payment", paymentHandler)
	http.ListenAndServe(":8080", nil)
}

func paymentHandler(w http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
	defer cancel()

	ch := make(chan bool)
	defer close(ch)

	// simulate payment processing
	go func() {
		fmt.Println("new payment processing...")
		sec := time.Duration(rand.Intn(15))

		select { // use select to wait for the context to be canceled or the payment to be processed
		case <-time.After(sec * time.Second): // simulate payment processing time between 0-15 seconds
			ch <- true // payment is successful
			fmt.Println("payment processed, after", sec)
		case <-ctx.Done():
			fmt.Printf("payment canceled after %v: %v\n", sec, ctx.Err())
			return // return to exit the goroutine
		}
	}()

	select {
	case result := <-ch: // if payment is successful, write response to the client
		if result {
			fmt.Fprintln(w, "payment successful")
			return
		}

		fmt.Fprintln(w, "payment failed")
	case <-ctx.Done(): // if the context is canceled before payment is done, write response to the client
		fmt.Fprintln(w, "payment canceled:", ctx.Err())
	}

	fmt.Println("payment handler ended")
}
