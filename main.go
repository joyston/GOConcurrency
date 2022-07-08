package main

import "fmt"

func swap(x int, c chan int) {
	c <- x * 2
}

func main() {
	c := make(chan int)
	go swap(2, c)
	go swap(3, c)
	y := <-c
	x := <-c
	fmt.Println(y, x)
}
