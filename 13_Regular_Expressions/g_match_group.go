package main

import (
	"fmt"
	"regexp"
)

func main() {
	// + minimum one time, or more times
	// () specifies a group
	re := regexp.MustCompile(`(abc)+`)
	fmt.Println(re.MatchString("abc"))    // true
	fmt.Println(re.MatchString("abcabc")) // true
	fmt.Println(re.MatchString("ab"))     // false
}
