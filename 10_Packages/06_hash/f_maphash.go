package main

import (
	"fmt"
	"hash/maphash"
)

// maphash is a non-cryptographic hash function
func main() {
	var h maphash.Hash
	h.WriteString("Hello, World!")
	fmt.Printf("MapHash: %x\n", h.Sum64())
}
