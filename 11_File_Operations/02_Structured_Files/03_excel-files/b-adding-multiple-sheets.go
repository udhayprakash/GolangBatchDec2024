package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()  // sheet1
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create multiple sheets
	f.NewSheet("Sheet2")
	f.NewSheet("Sheet3")

	// Set values in different sheets
	f.SetCellValue("Sheet1", "A10", "Sheet1 Data")
	f.SetCellValue("Sheet2", "A10", "Sheet2 Data")
	f.SetCellValue("Sheet3", "A10", "Sheet3 Data")

	// Save the file
	if err := f.SaveAs("MultipleSheets.xlsx"); err != nil {
		fmt.Println(err)
	}
}
