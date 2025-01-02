package main

import (
	"encoding/hex"
	"fmt"
	"reflect"
)

/*
hex Encoding

	10 Digits (0-9)
	 6 Uppercase letters (A-F) or lowercase letters (a-f)

	---
	16 bits -> 0 to 15


*/

func main() {

	orgMsg := "Hello world"
	fmt.Println("orgMsg=", orgMsg, reflect.TypeOf(orgMsg))

	byteStr := []byte(orgMsg)
	fmt.Println("byteStr=", byteStr, reflect.TypeOf(byteStr).Kind())

	// Encode
	encodedStr := hex.EncodeToString(byteStr)
	fmt.Println("encodedStr=", encodedStr, reflect.TypeOf(encodedStr).Kind())

	// Decode
	decodedMsg, err := hex.DecodeString(encodedStr)
	if err != nil {
		panic("malformed input. Unable to decode")
	}
	fmt.Println("decodedMsg=", decodedMsg, reflect.TypeOf(decodedMsg).Kind())

}

// Hello world
// SGVsbG8gd29ybGQ=					base64
// JBSWY3DPEB3W64TMMQ======			base32
// 48656c6c6f20776f726c64			hex


// assignment - try 	base64.RawURLEncoding  (url safe encoding)
