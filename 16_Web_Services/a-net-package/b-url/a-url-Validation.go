package main

import (
	"fmt"
	"net/url"
)

// isValidUrl tests a string to determine if it is a 
// well-structured url or not.
func isValidUrl(givenUrl string) bool {
	_, err := url.ParseRequestURI(givenUrl)
	if err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(isValidUrl("http://www.google.com")) // = true
	fmt.Println(isValidUrl("google.com"))            // = false
	fmt.Println(isValidUrl(""))                      // = false

	fmt.Println(isValidUrl("http://www.google.com/something")) // = true
	fmt.Println(isValidUrl("http://www.google.com/something?json=true")) // = true

}
