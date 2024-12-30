package main

import "fmt"

func main() {
	fmt.Println(divide(10, 0))
}

func divide(a, b int) int {
	return a / b
}

// troubleshoot zeroDivisionError panic

// b divide