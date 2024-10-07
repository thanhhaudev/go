package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- The worker pool is a design pattern used to manage concurrent tasks by creating a group of fixed workers to process tasks in parallel. Each worker is responsible for processing tasks sent to it.
- In Go:
  + The channel worker pool pattern creates a pool of goroutines that can process tasks concurrently.
  + The main goroutine sends tasks to a channel, and worker goroutines receive tasks from the channel and process them.
  + The number of worker goroutines in the pool is determined by the number of workers you create.
  + This pattern is useful for processing a large number of tasks concurrently.
  + It also helps to limit the number of goroutines created, providing better resource management.
- Pros:
  + Efficiently manages concurrent tasks.
  + Easy to implement, use, and maintain.
  + Limits the number of goroutines created, preventing resource exhaustion.
- Cons:
  + The number of workers in the pool is fixed; you must know the required number of workers in advance.
  + Not suitable for tasks that require dynamic scaling.
  + Difficult to manage complex dependencies between tasks or ensure task order.
*/

// each worker will block and wait for data to be sent to the channel.
// once the worker receives data, it processes the task and sends the result back to the main goroutine.
// the worker will continue to process tasks until the channel is closed.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs { // this loop will block and wait for data to be sent to the channel.
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}

	fmt.Printf("Worker %d stopped\n", id)
}

/*
- using wait group is not necessary, but it is a good practice to ensure that all workers have finished processing tasks before the main goroutine exits.
*/
func main() {
	const numJobs = 5
	const numWorkers = 3

	now := time.Now()
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		fmt.Println("Start worker", w)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		fmt.Println("Send job", j)
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	fmt.Println("Done after", time.Since(now))
}

/*
Result will be like this:
	Start worker 1
	Start worker 2
	Start worker 3
	Send job 1
	Send job 2
	Send job 3
	Send job 4
	Send job 5
	Worker 2 started job 2
	Worker 3 started job 3
	Worker 1 started job 1
	Worker 3 finished job 3
	Worker 1 finished job 1
	Worker 1 started job 5
	Worker 3 started job 4
	Worker 2 finished job 2
	Worker 2 stopped
	Worker 3 finished job 4
	Worker 3 stopped
	Worker 1 finished job 5
	Worker 1 stopped
	Result: 2
	Result: 4
	Result: 6
	Result: 8
	Result: 10
	Done after 2.00079241s
*/
