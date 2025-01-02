package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	data := []string{"apple", "banana", "apple", "orange", "banana"}
	seen := make(map[uint32]bool)

	for _, item := range data {
		hash := fnv.New32()
		hash.Write([]byte(item))
		if seen[hash.Sum32()] {
			fmt.Printf("Duplicate found: %s\n", item)
		} else {
			seen[hash.Sum32()] = true
		}
	}
}
