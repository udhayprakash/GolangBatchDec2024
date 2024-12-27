package main

import (
	"fmt"
	"reflect"
)

func main() {
	primesArray := [6]int{2, 3, 5, 7, 11, 13}
	// indexing -         0  1  2  3  4   5

	fmt.Printf("Type-%T   value-%[1]v\n", primesArray)
	fmt.Println(reflect.TypeOf(primesArray).Kind()) // array

	var slice1 []int = primesArray[1:4] // [3, 5, 7]
	fmt.Printf("Type-%T   value-%[1]v\n", slice1)
	fmt.Println(reflect.TypeOf(slice1).Kind()) // slice

	// mutability check
	fmt.Println("slice1[1]=", slice1[1]) // 5
	slice1[1] = 555

	// Changes in slices are reflected back to array
	fmt.Println("slice1     =", slice1)      // [3 555 7]
	fmt.Println("primesArray=", primesArray) // [2 3 555 7 11 13]

	fmt.Println("slice1[1:3]=", slice1[1:3])
	fmt.Println()

	// slice1[1:3] = []int{55555, 77777}
	// cannot assign to slice1[1:3]
	
}
