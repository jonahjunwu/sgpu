package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		r := rand.Intn(10)
		fmt.Printf("sleep %d\n", r)
		time.Sleep(time.Duration(r) * time.Second)
		results <- job * 2
	}
}

func main() {
	numJobs := 10
	numWorkers := 1

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}
	fmt.Print("enqueue jobs")
	// Enqueue jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	fmt.Print("wait for all workers to finish")
	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()
	fmt.Print("collect all result \n")
	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
