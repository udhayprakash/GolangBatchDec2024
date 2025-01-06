package main

import (
	"fmt"
	"time"
)

func longRunningOperation(stopchan chan struct{}) {
	for {
		select {
		case <-stopchan:
			fmt.Println("operation cancelled")
			return
		default:
			fmt.Println("Working .....")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stopchan := make(chan struct{})

	go longRunningOperation(stopchan)

	time.Sleep(2 * time.Second)
	close(stopchan) // cancel the operation
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting!!!")
}
