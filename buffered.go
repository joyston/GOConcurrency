package main

import "fmt"

func multiplAndSendOnChan(x int, c chan int) {
	c <- x * 2
}

func ExecuteBufferedChannels() {
	c := make(chan int)
	go multiplAndSendOnChan(2, c)
	go multiplAndSendOnChan(3, c)
	y := <-c
	x := <-c
	fmt.Println("Multiplied values:", y, x)
}
