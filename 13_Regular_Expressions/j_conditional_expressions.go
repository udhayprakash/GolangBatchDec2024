package main

import (
	"fmt"
	"regexp"
)

// Conditional Expressions (?(condition)true-pattern|false-pattern)

func main() {
	/*
		(a)?
			b(?(1)c
			d


			?=b(?=(1)c|d)
	*/
	re := regexp.MustCompile(`(a)?b|d`)
	fmt.Println(re.MatchString("abc")) // true
	fmt.Println(re.MatchString("bd"))  // true
	fmt.Println(re.MatchString("bc"))  // true
}
