package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)


// go get github.com/xuri/excelize/v2
func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create a new sheet
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		panic(err)
	}

	// Set value of a cell
	f.SetCellValue("Sheet2", "A2", "Hello world")
	f.SetCellValue("Sheet1", "B2", 100)

	// Set active sheet of the workbook
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
