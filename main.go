package main

import (
	"fmt"
	"time"
)

func multiplAndSendOnChan(x int, c chan int) {
	c <- x * 2
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := x; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func selectFibonacci(c, quit chan int) {
	a, b := 0, 1
	for {
		select {
		case c <- a:
			a, b = b, a+b
		case <-quit:
			fmt.Println("Quit Fibonacci using Select")
			return
		}
	}
}

func worker(done chan bool) {
	fmt.Println("......Started work")
	time.Sleep(time.Second)
	fmt.Println("work complete....")
	done <- true
}

func main() {
	cBool := make(chan bool)
	go worker(cBool)
	<-cBool //waiting to receive (Channel synchronization)

	c := make(chan int)
	go multiplAndSendOnChan(2, c)
	go multiplAndSendOnChan(3, c)
	y := <-c
	x := <-c
	fmt.Println("Multiplied values:", y, x)

	ch := make(chan int, 5)   //buffered channel
	go fibonacci(cap(ch), ch) //accept a limited number of values without a corresponding receiver for those values.

	for i := range ch {
		fmt.Println(i)
	}

	//Waiting on multiple channel operations using select
	cValues := make(chan int)
	cQuit := make(chan int)
	fmt.Println("********Fibonacci using Select*****")
	go func() {
		for j := 0; j < 10; j++ {
			fmt.Println(<-cValues)
		}
		cQuit <- 0
	}()
	selectFibonacci(cValues, cQuit)

	//Channel Directions
	ExecuteDirections()

	//Non-Blocking channels
	ExecuteNonBlocking()

	//Closing Channels
	ExecuteClosing()

	//Range over channels
	ExecuteRanging()
}
