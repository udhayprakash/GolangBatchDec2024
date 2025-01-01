package main

import (
	"fmt"
	"hash/adler32"
)

// Adler-32 is a checksum algorithm that is faster than CRC32 but less reliable.

func main() {
	data := []byte("Hello, World!")
	hash := adler32.Checksum(data)
	fmt.Printf("Adler-32 Checksum: %x\n", hash)
}
