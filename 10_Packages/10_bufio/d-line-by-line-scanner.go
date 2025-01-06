package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
