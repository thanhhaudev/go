package main

/*
Dynamic Worker Assignment: With chan chan Job, each worker can dynamically register itself to the workPool and wait for jobs. This allows the dispatcher to assign jobs to available workers dynamically.
Load Balancing: It helps in better load balancing. The dispatcher can pick any available worker from the pool and assign a job, ensuring that no worker is overloaded while others are idle.
Worker Reusability: Workers can be reused efficiently. Once a worker completes a job, it can re-register itself to the workPool and wait for the next job.
*/

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a unit of work to be processed by a worker.
type Job struct {
	ID   int
	Name string
}

// Worker represents a worker that processes jobs.
type Worker struct {
	ID       int
	Queues   chan Job
	WorkPool chan chan Job // Buffered channel containing unbuffered channels
	Stopped  chan bool
	wg       *sync.WaitGroup
}

// Start begins the worker's job processing loop.
// w.WorkPool <- w.Queues: Adds the worker's job queue to the workPool, signaling that the worker is ready to take on a job.
// select statement: Waits for either a job to be received on w.Queues or a stop signal on w.Stopped.
func (w *Worker) Start() {
	go func() {
		for { // the for {} loop ensures that the worker keeps running
			// Add the worker to the worker pool when idle, signaling that the worker is available to take on a job
			w.WorkPool <- w.Queues

			select {
			case job := <-w.Queues: // Receive a job from the worker's job queue
				fmt.Printf("🟡 Worker %d is processing job(%d): %s\n", w.ID, job.ID, job.Name)
				time.Sleep(2 * time.Second) // simulate work
				fmt.Printf("🟢 Worker %d has finished processing job(%d): %s\n", w.ID, job.ID, job.Name)
				w.wg.Done()
				// after job := <-w.Queues is done and the worker finishes processing the job, the worker will become idle again and wait for the next job.
			case <-w.Stopped:
				fmt.Printf("🔴 Worker %d is stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop signals the worker to stop processing jobs.
func (w *Worker) Stop() {
	w.Stopped <- true
	close(w.Queues) // Close the job queue to prevent further job assignments
}

// Dispatcher manages a pool of workers.
type Dispatcher struct {
	WorkPool   chan chan Job
	MaxWorkers int
	wg         *sync.WaitGroup
}

// Run starts the dispatcher and initializes workers.
func (d *Dispatcher) Run() {
	for i := 1; i <= d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkPool, d.wg)
		worker.Start() // Start the worker goroutine
	}

	fmt.Printf("Dispatcher is running with %d workers\n", d.MaxWorkers)
}

// dispatch sends a job to an available worker.
// Blocking Operations: Both jobQueue := <-d.WorkPool and jobQueue <- job are blocking operations. They ensure that a job is only dispatched when a worker is available and ready to process it.
// Synchronization: The sync.WaitGroup is used to synchronize the completion of jobs. The counter is incremented when a job is dispatched and decremented when a job is completed (handled in the worker's Start method).
// jobQueue := <-d.WorkPool will block until an available worker's job queue is received from the WorkPool channel. This ensures that a job is only dispatched when a worker is available.
// jobQueue <- job will block until the job is received by the worker's job queue. This ensures that the job is only sent when the worker is ready to process it.
func (d *Dispatcher) dispatch(job Job) {
	d.wg.Add(1)              // Increment the WaitGroup counter
	jobQueue := <-d.WorkPool // Get an available worker's job queue
	jobQueue <- job          // Send the job to the worker's job queue

	//fmt.Printf("Job(%d) has been dispatched\n", job.ID)
}

func (d *Dispatcher) Stop() {
	close(d.WorkPool) // Close the work pool to signal workers to stop
}

// NewDispatcher creates a new Dispatcher.
func NewDispatcher(maxWorkers int, wg *sync.WaitGroup) *Dispatcher {
	return &Dispatcher{
		WorkPool:   make(chan chan Job, maxWorkers),
		MaxWorkers: maxWorkers,
		wg:         wg,
	}
}

// NewWorker creates a new Worker.
func NewWorker(id int, workPool chan chan Job, wg *sync.WaitGroup) *Worker {
	return &Worker{
		ID:       id,
		Queues:   make(chan Job),
		Stopped:  make(chan bool),
		WorkPool: workPool,
		wg:       wg,
	}
}

func main() {
	var (
		now        = time.Now()
		wg         sync.WaitGroup
		numWorkers = 3
		numJobs    = 10
		namedJobs  = map[int]string{
			1:  "Clean the house",
			2:  "Wash the dishes",
			3:  "Do the laundry",
			4:  "Buy groceries",
			5:  "Cook dinner",
			6:  "Walk the dog",
			7:  "Do the gardening",
			8:  "Fold the clothes",
			9:  "Take out the trash",
			10: "Water the plants",
		}
	)

	dispatcher := NewDispatcher(numWorkers, &wg)
	dispatcher.Run()

	// Create and dispatch jobs
	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:   i,
			Name: namedJobs[i],
		}

		dispatcher.dispatch(job) // Dispatch the job
	}

	wg.Wait()         // Wait for all jobs to finish
	dispatcher.Stop() // Stop the dispatcher

	fmt.Println("All jobs are done, after:", time.Since(now))
}

/*
Run the program and observe the output:
	Dispatcher is running with 5 workers
	Worker 4 is processing job(4): Buy groceries
	Worker 2 is processing job(1): Clean the house
	Worker 1 is processing job(2): Wash the dishes
	Worker 5 is processing job(5): Cook dinner
	Worker 3 is processing job(3): Do the laundry
	Worker 3 is processing job(7): Do the gardening
	Worker 2 is processing job(8): Fold the clothes
	Worker 5 is processing job(6): Walk the dog
	Worker 1 is processing job(10): Water the plants
	Worker 4 is processing job(9): Take out the trash
	All jobs are done
*/
