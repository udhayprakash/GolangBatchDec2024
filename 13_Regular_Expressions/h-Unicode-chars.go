package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\p{Greek}`)
	fmt.Println(re.MatchString("α")) // true (Greek letter alpha)
	fmt.Println(re.MatchString("a")) // false
}
