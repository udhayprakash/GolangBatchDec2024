package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Method 1
	// \w
	re := regexp.MustCompile(`\w`)
	fmt.Println(re.MatchString("hello123")) // true
	fmt.Println(re.MatchString("!@#"))      // false

	match := re.MatchString("hello123")
	fmt.Println(match) // true

	// Method 2
	match, _ = regexp.MatchString(`\w`, "hello123")
	fmt.Println(match) // true
	fmt.Println()

	// w any charcter a-z , A-Z, 0-9, _
	re = regexp.MustCompile(`\w+`)
	fmt.Println(re.MatchString("hello123")) // true
	fmt.Println(re.MatchString("!@#"))      // false

}
