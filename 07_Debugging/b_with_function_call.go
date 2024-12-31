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


	> main.sum() ./b_with_function_call.go:12 (PC: 0x4af088)
	Command failed: expression *ast.CompositeLit not implemented
	
	(dlv) call nums[0] = 9
	> [Breakpoint 1] main.sum() ./b_with_function_call.go:12 (hits goroutine(1):1 

	(dlv) call nums = []{}int{9, 99, 999, 9999, 99999}
	> [Breakpoint 1] main.sum() ./b_with_function_call.go:12 (hits goroutine(1):1 total:1) (PC: 0x4af088)
	Command failed: 1:4: expected type, found '{'


	// not all operations can be done

	(dlv) nums = append(nums, 888)
	Command failed: command not available
	(dlv) call nums = append(nums, 888)
	xcr0: 0xe7 xstate_bv: 0x2
	> main.sum() ./b_with_function_call.go:12 (PC: 0x4af088)
	Command failed: could not find symbol value for append
	(dlv) set nums = append(nums, 888)
	Command failed: function calls not allowed without using 'call'
	*/


// dlv debug filename.go
