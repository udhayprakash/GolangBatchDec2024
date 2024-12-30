package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Goroutine:", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}

// dlv debug main.go
// break main.go:11
// continue
// goroutines // to list all active goroutines
// goroutine <id>  // switch to a specific goroutine
// step  // step through the goroutine
// quit  // Exit the debugger


// n line by line 
// step block by block 
// continue -- breakpoint wise 