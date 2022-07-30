package main

import "fmt"

func ping(cSending chan<- string, msg string) {
	cSending <- msg
}

func pong(cRec <-chan string, cSending chan<- string) {
	msg := <-cRec
	cSending <- msg
}

func ExecuteDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
