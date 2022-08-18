package main

import (
	"fmt"
	"time"
)

func ExecuteTickers() {
	ticker1 := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker1.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker1.Stop()
	done <- true
	fmt.Println("Ticker  Stopped")
}
