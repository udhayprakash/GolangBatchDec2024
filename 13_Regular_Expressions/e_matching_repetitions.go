package main

import (
	"fmt"
	"regexp"
)

func main() {
	// exact repeitions , to match
	re := regexp.MustCompile(`a{2}`)
	fmt.Println(re.MatchString("aa"))  // true
	fmt.Println(re.MatchString("a"))   // false
	fmt.Println(re.MatchString("aaa")) // true

	// min and max repition count , to match
	re = regexp.MustCompile(`a{2,4}`)
	fmt.Println(re.MatchString("a"))     // false
	fmt.Println(re.MatchString("aa"))    // true
	fmt.Println(re.MatchString("aaa"))   // true
	fmt.Println(re.MatchString("aaaa"))  // true
	fmt.Println(re.MatchString("aaaaa")) // true (4 present)
}
