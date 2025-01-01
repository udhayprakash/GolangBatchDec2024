package main

import (
	"encoding/base32"
	"fmt"
	"reflect"
)

/*
Base32 Encoding

	26 Uppercase letters (A-Z)
	9 Digits (0-7)
	---
	32 bits


	binary data --> groups of 5 bits -> each group is mapped to base32 character
	if input  data is not a multiple of 5 bytes, padding char (=) added in the end

*/

func main() {

	orgMsg := "Hello world"
	fmt.Println("orgMsg=", orgMsg, reflect.TypeOf(orgMsg))

	byteStr := []byte(orgMsg)
	fmt.Println("byteStr=", byteStr, reflect.TypeOf(byteStr).Kind())

	// Encode
	encodedStr := base32.StdEncoding.EncodeToString(byteStr)
	fmt.Println("encodedStr=", encodedStr, reflect.TypeOf(encodedStr).Kind())

	// Decode
	decodedMsg, err := base32.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic("malformed input. Unable to decode")
	}
	fmt.Println("decodedMsg=", decodedMsg, reflect.TypeOf(decodedMsg).Kind())

}

// Hello world
// SGVsbG8gd29ybGQ=					base64
// JBSWY3DPEB3W64TMMQ======			base32
