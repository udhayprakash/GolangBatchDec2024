package main

import (
	as "calculator/addsub"
	. "calculator/muldiv"
	"fmt"
)

func main() {
	fmt.Println("main() function start")

	// fmt.Println(os.Getwd())

	resultAdd := as.Addition(12, 10)
	fmt.Println("resultAdd =", resultAdd)

	fmt.Println("as.Subtraction(12, 10) =", as.Subtraction(12, 10))

	fmt.Println(Divide(10, 3) == DivideInts(10, 3))

	// multiply(10, 3)

	fmt.Println("Filename = ", Filename)
	fmt.Println("as.Filename = ", as.Filename)

}
