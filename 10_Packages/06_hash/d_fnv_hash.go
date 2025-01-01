package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	data := []byte("Hello, World!")

	// FNV-1 (32-bit)
	hash := fnv.New32()
	hash.Write(data)
	fmt.Printf("FNV-1 (32-bit) Hash: %x\n", hash.Sum32())

	// FNV-1a (64-bit)

	hash2 := fnv.New64a()
	hash.Write(data)
	fmt.Printf("FNV-1a (64-bit) Hash: %x\n", hash2.Sum64())

}
