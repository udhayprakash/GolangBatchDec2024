package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`https?://[^\s]+`)
	text := "Visit https://example.com and http://another.com"
	urls := re.FindAllString(text, -1)
	fmt.Println(urls) // [https://example.com http://another.com]
}
