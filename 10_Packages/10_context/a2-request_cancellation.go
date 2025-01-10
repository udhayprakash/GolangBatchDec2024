// request cancellation, with context
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation2(ct context.Context, jobName string) {
	for {
		select {
		case <-ct.Done():
			fmt.Printf("[%s] Operation cancelled: %v\n", jobName, ct.Err())
			return
		default:
			fmt.Printf("[%s] Working...\n", jobName)
			time.Sleep(500 * time.Millisecond)

		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go longRunningOperation2(ctx, "job1")
	go longRunningOperation2(ctx, "job2")

	time.Sleep(2 * time.Second)
	cancel() // Cancel the operation
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting.")

}
