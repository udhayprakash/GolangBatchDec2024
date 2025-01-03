package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Set values
	f.SetCellValue("Sheet1", "A1", 10)
	f.SetCellValue("Sheet1", "B1", 20)

	// Set formula
	f.SetCellFormula("Sheet1", "C1", "=A1+B1")

	// Save the file
	if err := f.SaveAs("CellReferenceFormula.xlsx"); err != nil {
		fmt.Println(err)
	}
}
