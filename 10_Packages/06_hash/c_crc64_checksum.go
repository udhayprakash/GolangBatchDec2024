package main

import (
	"fmt"
	"hash/crc64"
)

// 64 bit hash value

func main() {
	data := []byte("Hello, World!")
	table := crc64.MakeTable(crc64.ISO)
	hash := crc64.Checksum(data, table)
	fmt.Printf("CRC64 Checksum: %x\n", hash)
}