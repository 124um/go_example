package main

import (
	"fmt"
	"time"
)

func mainWork() {

	t := time.Now()

	const jobsCount, workerCount int = 40, 20
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	for i := 0; i < workerCount; i++ {
		go worker(i+1, jobs, results)

	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i + 1
	}
	close(jobs)

	for i := 0; i < jobsCount; i++ {
		fmt.Printf("result #%d : value %d \n", i+1, <-results)
	}

	fmt.Println("Time for work - ", time.Since(t).String())
}

// jobs <-chan int    - только читать с канала
// results chan<- int   - только писать в канал

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("worker #%d finished\n", id)
		results <- j * j
	}
}