package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

var (
	filepath = `D:\Test\Consumer\Benchmark AI article review_batch25_20200403.xlsx`
)

func main() {
	xlfile, err := xlsx.OpenFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, sheet := range xlfile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		fmt.Println(sheet.MaxRow)
		sheet.Cell(5, 5)
		fmt.Println(sheet.Cell(1, 1))
	}

}
