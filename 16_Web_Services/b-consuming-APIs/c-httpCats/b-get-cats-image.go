package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	url := "https://http.cat/200"

	// Fetch the HTML content
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the page:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 (OK)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Received non-200 status code:", resp.StatusCode)
		return
	}

	file, err := os.Create("http_cat_200.jpg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// copying from buffer to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return
	}

	fmt.Println("Image downloaded successfully: http_cat_200.jpg")

}
