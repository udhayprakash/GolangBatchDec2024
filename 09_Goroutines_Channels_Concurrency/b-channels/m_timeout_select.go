package main

/*
Purpose: Timeouts
	- are important for programs that connect to external
	  resources or that otherwise need to bound execution time.
	- Implementing timeouts in Go is easy and elegant thanks
      to channels and select.
*/
import (
	"fmt"
	"time"
)

func main() {
	// Scenario 1
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Scenario 2
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}