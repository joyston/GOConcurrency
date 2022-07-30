package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("......Started work")
	time.Sleep(time.Second)
	fmt.Println("work complete....")
	done <- true
}

func ExecuteSynchronize() {
	cBool := make(chan bool)
	go worker(cBool)
	<-cBool //waiting to receive (Channel synchronization)
}
