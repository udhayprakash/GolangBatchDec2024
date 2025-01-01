package main

import (
	"fmt"
	"hash/fnv"
)

type HashTable struct {
	data map[uint32]string
}

func NewHashTable() *HashTable {
	return &HashTable{data: make(map[uint32]string)}
}

func (h *HashTable) Set(key, value string) {
	hash := fnv.New32()
	hash.Write([]byte(key))
	h.data[hash.Sum32()] = value
}

func (h *HashTable) Get(key string) (string, bool) {
	hash := fnv.New32()
	hash.Write([]byte(key))
	value, exists := h.data[hash.Sum32()]
	return value, exists
}

func main() {
	table := NewHashTable()
	table.Set("name", "Alice")
	table.Set("age", "30")

	fmt.Println(table.Get("name")) // Output: Alice true
	fmt.Println(table.Get("age"))  // Output: 30 true
	fmt.Println(table.Get("city")) // Output:  false
}
