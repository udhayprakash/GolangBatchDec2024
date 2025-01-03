package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() != false {
			fmt.Println(path, info.Name())
		}
		return nil
	})
}

// Assignment:
//  count the number of differenr types files in the current project repo level
// {"go": 0, "text": 2}
