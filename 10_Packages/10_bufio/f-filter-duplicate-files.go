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

	// Create a map to store line counts
	lineCounts := make(map[string]int)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		lineCounts[line]++
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(lineCounts)

	// Print duplicate lines
	for line, count := range lineCounts {
		if count > 1 {
			fmt.Printf("Duplicate line: %s (count: %d)\n", line, count)
		}
	}
}
