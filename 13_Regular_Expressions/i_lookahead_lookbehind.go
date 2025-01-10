package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Positive lookahead
	re := regexp.MustCompile(`foo(?=bar)`)
	fmt.Println(re.MatchString("foobar")) // true
	fmt.Println(re.MatchString("foobaz")) // false

	// Negative lookahead
	re = regexp.MustCompile(`foo(?!bar)`)
	fmt.Println(re.MatchString("foobar")) // false
	fmt.Println(re.MatchString("foobaz")) // true

	// Positive lookbehind
	re = regexp.MustCompile(`(?<=foo)bar`)
	fmt.Println(re.MatchString("foobar")) // true
	fmt.Println(re.MatchString("bazbar")) // false

	// Negative lookbehind
	re = regexp.MustCompile(`(?<!foo)bar`)
	fmt.Println(re.MatchString("foobar")) // false
	fmt.Println(re.MatchString("bazbar")) // true

}
