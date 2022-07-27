package main

import "fmt"

func ExecuteNonBlocking() {

	cRec := make(chan string)
	cSend := make(chan string)

	//non-blocking receive
	select {
	case msg := <-cRec:
		fmt.Println("Received: ", msg)
	default:
		fmt.Println("No msg Received")
	}

	//non-blocking send
	msg := "hi"
	select {
	case cSend <- msg:
		fmt.Println("Sent: ", msg)
	default:
		fmt.Println("No msg Sent")
	}

	//multiple cases
	cBool := make(chan bool)
	select {
	case msg := <-cRec:
		fmt.Println("String received", msg)
	case lbool := <-cBool:
		fmt.Println("Boolean received", lbool)
	case cSend <- msg:
		fmt.Println("String sent", msg)
	default:
		fmt.Println("Nothing received/Sent")
	}
}
