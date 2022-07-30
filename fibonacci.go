package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := x; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func ExecuteFibonacci() {
	ch := make(chan int, 5)   //buffered channel
	go fibonacci(cap(ch), ch) //accept a limited number of values without a corresponding receiver for those values.

	for i := range ch {
		fmt.Println(i)
	}
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

func ExecuteSelectFibonacci() {
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

}
