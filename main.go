package main

import "fmt"

func swap(x int, c chan int) {
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

func main() {
	c := make(chan int)
	go swap(2, c)
	go swap(3, c)
	y := <-c
	x := <-c
	fmt.Println(y, x)

	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}
}