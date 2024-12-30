// partial function
package main

import "fmt"

/*
2 * 3 = 6 	=>  multiply(2, 3)
2 * 4 = 8  =>  multiply(2, 4)
2 * 9 = 18 =>  multiply(2, 9)
*/
func multiply(a, b int) int {
	return a * b
}

// Higher Order Function
func multiplyWithTwo(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}


// Higher order function for adding three numbers
func adder(x, y int) func(int) int {
	return func(z int) int {
		return x + y + z
	}
}


func main() {
	fmt.Println(multiply(2, 3))
	fmt.Println(multiply(2, 4))
	fmt.Println(multiply(2, 9))
	fmt.Println()

	fmt.Println(multiplyWithTwo(2)(3))

	mul2 := multiplyWithTwo(2)
	fmt.Println(mul2(4))
	fmt.Println(mul2(9))
	fmt.Println()

	partialAdd := adder(3, 4)
	fmt.Println("partialAdd(5) = ", partialAdd(5))
	fmt.Println("partialAdd(6) = ", partialAdd(6))

}


// Assignment - create parital function for building sql query , to filter by name
