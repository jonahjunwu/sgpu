package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/exec"
	"sync"
	"time"

	"github.com/deliveryhero/pipeline/v2"
)

func gpurun(wisperJob string) {
	pipeInitial := pipeline.NewProcessor(func(_ context.Context, s []string) ([]string, error) {
		fmt.Println("pip initial check", s)
		return s, nil
		//return strings.Split(s, ","), nil
	}, nil)

	pipeP1 := pipeline.NewProcessor(func(_ context.Context, s string) (string, error) {
		cmd1 := exec.Command("python3", "p1python.py")
		out1, _ := cmd1.CombinedOutput()
		fmt.Printf("output: %s", out1)
		fmt.Println("P1 working on", wisperJob)
		fmt.Println("P1 recieve", s)
		return "P1 finished", nil
	}, nil)

	pipeP2 := pipeline.NewProcessor(func(_ context.Context, s string) (string, error) {
		cmd1 := exec.Command("python3", "p2python.py")
		out1, _ := cmd1.CombinedOutput()
		fmt.Printf("output: %s", out1)
		fmt.Println("P2 working on", wisperJob)
		fmt.Println("P2 recieve", s)
		return "P2 finished", nil
	}, nil)

	pipeP3 := pipeline.NewProcessor(func(_ context.Context, s string) (string, error) {
		cmd1 := exec.Command("python3", "p3python.py")
		out1, _ := cmd1.CombinedOutput()
		fmt.Printf("output: %s", out1)
		fmt.Println("P3 working on", wisperJob)
		fmt.Println("P3 recieve", s)
		fmt.Println(s)
		return "P3 finished", nil
	}, nil)

	apply := pipeline.Apply(
		pipeInitial,
		pipeline.Sequence(
			pipeP1,
			pipeP2,
			pipeP3,
		),
	)
	//var input = []string{"1"}
	var input = []string{wisperJob}
	//input := "1"

	for out := range pipeline.Process(context.Background(), apply, pipeline.Emit(input)) {
		for j := range out {
			fmt.Printf("process: %s\n", out[j])
		}
	}
}

func fsleep() {
	r := rand.Intn(10)
	fmt.Printf("sleep %d\n", r)
	time.Sleep(time.Duration(r) * time.Second)
}

func worker(id int, jobs <-chan string, results chan<- string) {
	fmt.Println("workerid:", id)
	for job := range jobs {
		fmt.Println("start new job %d", job)
		fmt.Println("Worker", id, "processing job", job)
		fsleep()
		gpurun(job)
		results <- job
	}
}

func main() {
	numWorkers := 3
	whisperFiles, _ := readJsonFile("input.json")
	fmt.Println("read whisper files from file:", whisperFiles)
	//whisperFiles := []string{"1111", "2222", "3333", "4444", "5555", "6666", "7777", "8888", "9999", "1010101"}
	numJobs := len(whisperFiles)
	//numJobs := 10

	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)

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
		jobs <- whisperFiles[i-1]
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
		fmt.Println("Result:", result)
	}
}
