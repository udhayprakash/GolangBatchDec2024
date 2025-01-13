package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var foldername = "httpCatImages"

func init() {
	err := os.MkdirAll("httpCatImages", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}

func main() {
	baseURL := "https://http.cat/"

	for statusCode := 100; statusCode <= 599; statusCode++ {
		imageURL := baseURL + strconv.Itoa(statusCode)

		// Fetch the HTML content
		resp, err := http.Get(imageURL)
		if err != nil {
			fmt.Println("Error fetching the page:", err)
			return
		}
		defer resp.Body.Close()

		// Check if the response status code is 200 (OK)
		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("Image not found for this code", statusCode)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf(imageURL)
			fmt.Println("Error: Received non-200 status code:", resp.StatusCode)

			return
		}

		imageFileName := fmt.Sprintf("%s/%d.jpg", foldername, statusCode)
		file, err := os.Create(imageFileName)
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

	}
	fmt.Println("all Image downloaded successfully")

}
