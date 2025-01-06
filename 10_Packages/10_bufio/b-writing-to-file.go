package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a new buffered writer
	writer := bufio.NewWriter(file)

	// write data
	writer.WriteString("Hello world")
	res, err := writer.WriteString("This is second line")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println(res) // 19

	_, err = writer.WriteString("This is third line")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// flush , to ensure all data is written
	writer.Flush()

}
