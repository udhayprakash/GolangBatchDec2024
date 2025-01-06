package main

import (
	"fmt"
	"regexp"
)

func main() {
	// match a
	re := regexp.MustCompile("a")         // pattern to search
	fmt.Println(re.MatchString("apple"))  // true
	fmt.Println(re.MatchString("bcdfg"))  // false
	fmt.Println(re.MatchString("bcdfgu")) // true

	// match vowels
	re = regexp.MustCompile("[aeiou]")    // pattern to search
	fmt.Println(re.MatchString("apple"))  // true
	fmt.Println(re.MatchString("bcdfg"))  // false
	fmt.Println(re.MatchString("bcdfgu")) // true

	// match lower case letters
	re = regexp.MustCompile("[a-z]")      // pattern to search
	fmt.Println(re.MatchString("apple"))  // true
	fmt.Println(re.MatchString("bcdfg"))  // false
	fmt.Println(re.MatchString("bcdfgu")) // true

	// match digits
	re = regexp.MustCompile(`\d`)             // pattern to search; observe the quotes
	fmt.Println(re.MatchString("12312312"))   // true
	fmt.Println(re.MatchString("bcdfg"))      // false
	fmt.Println(re.MatchString("bcdfgu1231")) // true
}
