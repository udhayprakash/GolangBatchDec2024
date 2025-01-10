package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// Create an io.Reader (in this case, a string reader)
	reader := strings.NewReader("Hello, World!\nThis is a test.\n")

	// Create a new buffered reader
	bufReader := bufio.NewReader(reader)

	// Read the data
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}
