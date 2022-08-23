package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d Completed\n", id)
}

func ExecWaitGroups() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i //to avoid using the same variable for the goroutine
		go func() {
			defer wg.Done()
			task(i)
		}()
	}
	wg.Wait() //To wait for multiple goroutines to finish
}
