package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	hash := crc32.NewIEEE()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("File Checksum: %x\n", hash.Sum32())
}
