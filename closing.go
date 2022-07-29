package main

import "fmt"

func ExecuteClosing() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recieved job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i < 4; i++ {
		jobs <- i
		fmt.Println("Sent job", i)
	}
	close(jobs) //Closing channel after all 3 jobs are sent
	fmt.Println("All jobs sent")
	<-done //Code to synchronise and wait

}
