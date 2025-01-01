package main

import (
	"encoding/base64"
	"fmt"
	"reflect"
)

/*
Base64 Encoding

	26 Uppercase letters (A-Z)
	26 Lowercase letters (a-z)
	10 Digits (0-9)
	2  Two additional characters (+ and /)
	---
	64 bits


	binary data --> groups of 6 bits -> each group is mapped to base64 character
	if input  data is not a multiple of 3 bytes, padding char (=) added in the end

*/

func main() {

	orgMsg := "Hello world"
	fmt.Println("orgMsg=", orgMsg, reflect.TypeOf(orgMsg))

	byteStr := []byte(orgMsg)
	fmt.Println("byteStr=", byteStr, reflect.TypeOf(byteStr).Kind())

	// Encode
	encodedStr := base64.StdEncoding.EncodeToString(byteStr)
	fmt.Println("encodedStr=", encodedStr, reflect.TypeOf(encodedStr).Kind())

	// Decode
	decodedMsg, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic("malformed input. Unable to decode")
	}
	fmt.Println("decodedMsg=", decodedMsg, reflect.TypeOf(decodedMsg).Kind())

}


// Hello world 
// SGVsbG8gd29ybGQ=