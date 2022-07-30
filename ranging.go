package main

import "fmt"

func ExecuteRanging() {
	cRange := make(chan string, 2)

	cRange <- "Sending 1"
	cRange <- "Sending 2"
	close(cRange)

	for ltemp := range cRange {
		fmt.Println(ltemp)
	}
}
