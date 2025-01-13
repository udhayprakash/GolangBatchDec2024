package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passwd", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.SetBasicAuth("user", "passwd")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(body))
}
