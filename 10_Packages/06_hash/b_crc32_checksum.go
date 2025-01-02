package main

import (
	"fmt"
	"hash/crc32"
)

// CRC32 is commonly used for error-checking in network protocols and file integrity verification.

func main() {
	data := []byte("Hello, World!")
	hash := crc32.ChecksumIEEE(data)
	fmt.Printf("CRC32 Checksum: %x\n", hash)
}
