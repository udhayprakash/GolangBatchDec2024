package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://http.cat/"

	// Fetch the HTML content
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the page:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(body))
}
