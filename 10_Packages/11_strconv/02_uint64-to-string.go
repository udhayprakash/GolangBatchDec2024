package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num1 int = 99
	fmt.Println(num1, string(num1)) // 99  c

	// Create our number
	var myNumber uint64
	myNumber = 18446744073709551615

	str := strconv.FormatUint(myNumber, 10)
	fmt.Println("The number is: " + str)

}
