package main

import (
	"fmt"
	"regexp"
)

func main() {
	// if you have regex character in code, you need to escape
	// escape singe char [.]  or \.
	re := regexp.MustCompile(`[a-zA-Z._0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+\.?[a-zA-Z0-9]+`)
	fmt.Println(re.MatchString("test@example.com"))   // true
	fmt.Println(re.MatchString("test@example.co.uk")) // true
	fmt.Println(re.MatchString("test@example.co.in")) // true
	fmt.Println(re.MatchString("test@examplecom"))    // false
	fmt.Println(re.MatchString("invalid-email"))      // false

	// \w  [a-zA-Z0-9_]
	re = regexp.MustCompile(`\w.+@\w+\.\w+\.?\w+`)
	fmt.Println(re.MatchString("test@example.com"))   // true
	fmt.Println(re.MatchString("test@example.co.uk")) // true
	fmt.Println(re.MatchString("test@example.co.in")) // true
	fmt.Println(re.MatchString("test@examplecom"))    // false
	fmt.Println(re.MatchString("invalid-email"))      // false

}

// Assignment:
// 1. FInd all the valid Phone numbers in the given target string
