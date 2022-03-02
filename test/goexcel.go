package main

import (
	"fmt"
	"math/rand"

	"github.com/xuri/excelize/v2"
)

func main() {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")

	streamWriter.MergeCell("A1", "A2")
	if err != nil {
		fmt.Println(err)
	}
	err = streamWriter.SetColWidth(1, 1, 15)
	styleID, err := file.NewStyle(`{"font":{"bold":true},"fill":{"type":"pattern","color":["#808080"],"pattern":1}}`)
	if err != nil {
		fmt.Println(err)
	}
	ColTitle := make([]interface{}, 50)
	for colID := 0; colID < 50; colID++ {
		ColTitle[colID] = excelize.Cell{StyleID: styleID, Value: rand.Intn(640000)}
	}

	err = streamWriter.SetRow("A1", ColTitle)

	cell, _ := excelize.CoordinatesToCellName(1, 1)
	if err := streamWriter.SetRow(cell, []interface{}{
		excelize.Cell{StyleID: styleID, Value: []string{"Data", "Data1"}}}); err != nil {
		fmt.Println(err)
	}
	for rowID := 3; rowID <= 1024; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}

	err = file.MergeCell("Sheet1", "D3", "E9")
	fmt.Println("err:", err)
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
