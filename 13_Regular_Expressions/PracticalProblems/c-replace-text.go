package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\bfoo\b`)
	text := "foo bar foo baz"
	newText := re.ReplaceAllString(text, "bar")
	fmt.Println(newText) // "bar bar bar baz"
}
