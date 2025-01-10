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

    // Set formula
    f.SetCellFormula("Sheet1", "A3", "=SUM(A1:A2)")

    // Save the file
    if err := f.SaveAs("FormulaExample.xlsx"); err != nil {
        fmt.Println(err)
    }
}
