package main

import (
	"fmt"
	"time"
)

func workerpool(id int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Println("Worker ", id, " started job ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker ", id, "finished job", j)
		result <- j * 2
	}
}

func ExecWorkerPool() {
	const numJobs = 5

	job := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go workerpool(i, job, result)
	}

	for j := 1; j <= numJobs; j++ {
		job <- j
	}
	close(job)

	for a := 1; a <= numJobs; a++ {
		<-result
	}
}
