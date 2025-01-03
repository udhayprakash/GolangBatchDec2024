package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Person struct {
	Name     string
	Age      int
	Location string
	country  string
}

func main() {
	// Open the CSV file
	file, err := os.Open("personDetails.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for i, record := range records {
		if i == 0 { // Skip the header
			continue
		}

		fmt.Println(record[1])
	}

}
