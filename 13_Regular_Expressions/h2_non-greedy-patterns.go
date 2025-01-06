package main

import (
	"fmt"
	"regexp"
)
// ?P -- Perl compatible
func main() {
	re := regexp.MustCompile(`(?P<first>\w+) (?P<last>\w+)`)
	matches := re.FindStringSubmatch("John Doe")
	fmt.Println(matches) // [John Doe John Doe]
	fmt.Println(re.SubexpNames()) // [ first last]
}