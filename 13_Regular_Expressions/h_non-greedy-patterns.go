package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a`)
	fmt.Println(re.FindString("aaaa")) // "a"

	re = regexp.MustCompile(`a+`)
	fmt.Println(re.FindString("aaaa")) // "aaaa"

	// ? -- 0 or 1 time only
	re = regexp.MustCompile(`a+?`)
	fmt.Println(re.FindString("bbb"))  // ""
	fmt.Println(re.FindString("aaaa")) // "a"

	re = regexp.MustCompile(`<h1>my text</h1>`)
	fmt.Println(re.FindString("<.+>")) //
	fmt.Println(re.FindString("<.+?>")) //
}
