package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/stream/5")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// direct
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println(string(body))
	fmt.Println()
	
	// stream consumption
	buf := make([]byte, 1024) // reading one byte, at a time
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading stream:", err)
			return
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
