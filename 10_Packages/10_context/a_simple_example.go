package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	doWork(ctx)

}

func doWork(ctx context.Context) {

	for i := 0; i < 5; i++ { // 0 to 4
		select {
		case <-ctx.Done():
			fmt.Println("workd canceled")
			return
		default:
			fmt.Printf("working: %d\n", i)
			time.Sleep(2 * time.Second)
		}

	}
	fmt.Println("word completd")
}
