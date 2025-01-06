package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\s+`)
	text := "split   this   text"
	parts := re.Split(text, -1)
	fmt.Println(parts) // [split this text]
}
