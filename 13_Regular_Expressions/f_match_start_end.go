package main

import (
	"fmt"
	"regexp"
)

func main() {
	// ^ (carrot operator)-- at the start of string
	re := regexp.MustCompile(`^hello`)
	fmt.Println(re.MatchString("hello world")) // true
	fmt.Println(re.MatchString("world hello")) // false

	// $ (dollor operator)-- at the end of string
	re = regexp.MustCompile(`world$`)
	fmt.Println(re.MatchString("hello world")) // true
	fmt.Println(re.MatchString("world hello")) // false
}
