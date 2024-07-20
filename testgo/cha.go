package main

import "fmt"

func main() {
	numJobs := 10
	jobs := make(chan int, numJobs)
	fmt.Println(jobs)
}
