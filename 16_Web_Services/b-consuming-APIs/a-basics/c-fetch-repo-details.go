package main

import (
	"log"
	"net/http"
)

func main() {
	url := "https://api.github.com/repos/udhayprakash/GolangBatchDec2024"

	// // http.Get()
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal("Error creating request:", err)
	// }
	// defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(body))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "Go-CRUD-Example")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()
}
