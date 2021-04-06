package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func readexcle() {
	excelFileName := "test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("Sheet Name: %s\n", sheet.Name)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}

type struceTest struct {
	Name  string `xlsx: ""`
	Age   int    `xlsx:"20"`
	Score int    `xlsx:"0"`
}

func WriteExcel() {
	structVal := struceTest{
		Name:  "ken",
		Age:   18,
		Score: 99,
	}
	f, _ := xlsx.OpenFile("test1.xlsx")
	// sheet, _ := f.AddSheet("TestRead")
	sheet := f.Sheet["TestRead"]
	row := sheet.AddRow()
	row.WriteStruct(&structVal, -1)
	f.Save("test1.xlsx")
}
func main() {
	WriteExcel()
}
