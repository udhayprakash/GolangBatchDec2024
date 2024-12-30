package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("Sum of numbers:", sum(numbers))
}

func sum(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

/*
start debugger 

	dlv debug main.go

To show current executing line 
	list

set breakpoint

	break sum
	break b_with_function_call.go:10
	break b_with_function_call.go:13
	break b_with_function_call.go:15

List all breakpoints

	breakpoints

continue to next steps 

	continue


clear breakpoints
	clear 1		// to clear assigned breakpoints, based on order of assignment (not line number), from 1
	clearall    // to clear all breakpoints

Step through the loop:

	step

Exit the debugger

	quit
*/


// dlv debug filename.go
