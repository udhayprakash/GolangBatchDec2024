package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("hello")          // pattern to search
	fmt.Println(re.MatchString("hello world")) // true
	fmt.Println(re.MatchString("hell o"))      // false
	fmt.Println(re.MatchString("hi "))         // false

	fmt.Println(re.MatchString(`
		hello
		hello world
		hi
		hell o
	
	`))
}
