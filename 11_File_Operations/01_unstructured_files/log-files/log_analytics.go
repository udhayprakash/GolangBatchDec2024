package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func prettyPrintMap(m map[string]int) {
	fmt.Println("--------")
	fmt.Println("KEY\tVALUE")
	fmt.Println("--------")
	for k, v := range m {
		fmt.Printf("%-8s %d\n", k+":", v)
	}
	fmt.Println("--------")
}

func main() {
	// Open the access.log file for reading
	file, err := os.Open("access.log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file line by line and count the URLs
	scanner := bufio.NewScanner(file)

	mostAccessedURL, mostAccessedCount := "", 0
	urlCounts := make(map[string]int)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		// fmt.Println(len(fields), fields)
		if len(fields) <= 5 {
			continue
		}
		url := strings.Split(fields[6], "?")[0]
		urlCounts[url]++

		if urlCounts[url] > mostAccessedCount {
			mostAccessedURL, mostAccessedCount = url, urlCounts[url]

		}

	}
	fmt.Println(urlCounts)
	prettyPrintMap(urlCounts)

	// Print the result
	fmt.Printf("The most accessed URL is %s with %d hits.\n", mostAccessedURL, mostAccessedCount)
}
