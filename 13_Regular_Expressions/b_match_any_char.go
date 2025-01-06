package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("h.llo")          // pattern to search
	fmt.Println(re.MatchString("hello world")) // true
	fmt.Println(re.MatchString("hell o"))      // false
	fmt.Println(re.MatchString("hi "))         // false

	fmt.Println(re.MatchString("hallo")) // true
	fmt.Println(re.MatchString("h9llo")) // true
	fmt.Println(re.MatchString("h&llo")) // true
	fmt.Println(re.MatchString("h.llo")) // true

}
