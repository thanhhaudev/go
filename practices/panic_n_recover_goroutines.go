package main

/*
- A sync.WaitGroup in Go is used to wait for a collection of goroutines to finish executing.
It provides a way to synchronize the completion of multiple goroutines.
When you call wg.Wait(), it blocks the execution of the current goroutine until the counter inside the WaitGroup becomes zero,
indicating that all goroutines have finished their work.

*/
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go safeRoutine()
	go safeRoutine()
	time.Sleep(1 * time.Second)
	// The time.Sleep(1 * time.Second) call is used to pause the main goroutine for 1 second, allowing the other goroutines (safeRoutine) to execute.
	// Without this sleep, the main goroutine might exit before the other goroutines have a chance to run, which would terminate the program prematurely.

	// Is 1 second enough time for the goroutines to finish?
	// The duration of time.Sleep(1 * time.Second) is generally sufficient for simple examples where goroutines are expected to complete quickly.
	// However, in more complex scenarios, you may need to adjust the duration based on the expected execution time of the goroutines.
	// If the duration is too short, the main goroutine may exit before the other goroutines have finished, causing the program to terminate prematurely.
	// If the duration is too long, it may introduce unnecessary delays in the program execution.
	// In practice, you should adjust the sleep duration based on the specific requirements of your program.

	fmt.Println("This line will be executed.")

	// instead of using time.Sleep, we can use sync.WaitGroup to wait for the goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(2) // add 2 goroutines to the WaitGroup
	go func() {
		defer wg.Done() // defer Done() to indicate that the goroutine has finished
		// if the goroutine panics, Done() will still be called?
		// Yes, if the goroutine panics, the Done() method will still be called because it is deferred.
		safeRoutine()
	}()
	go func() {
		defer wg.Done() // The error panic: sync: negative WaitGroup counter occurs when the Done method is called more times than the Add method
		safeRoutine()
	}()

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines have finished.")
}

func safeRoutine() {
	defer func() { // The defer statement ensures that the Done() method is executed when the goroutine exits regardless of whether it exits normally or due to a panic
		if r := recover(); r != nil { // if we don't recover from the panic, the program will crash.
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Panic!")
}
