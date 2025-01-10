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
	f.SetCellValue("Sheet1", "A2", 20)
	f.SetCellValue("Sheet1", "A3", 30)

	// Set formula
	f.SetCellFormula("Sheet1", "A4", "=AVERAGE(A1:A3)")

	// Save the file
	if err := f.SaveAs("FunctionFormula.xlsx"); err != nil {
		fmt.Println(err)
	}
}
