package main

import "fmt"

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var num int
	return func() int {
		num++
		return num * num
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9

	f2 := squares()
	fmt.Println(f2()) // 1
	fmt.Println(f2()) // 4
}