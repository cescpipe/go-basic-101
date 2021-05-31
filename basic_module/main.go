package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
)

func AddCell(sheet *xlsx.Sheet, row, col int) *xlsx.Cell {
	for row >= len(sheet.Rows) {
		sheet.AddRow()
	}
	for col >= len(sheet.Rows[row].Cells) {
		sheet.Rows[row].AddCell()
	}
	return sheet.Cell(row, col)
}

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error


	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil{
		log.Fatal("error new sheet", err)
	}

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "I am a cell!"
	cell = AddCell(sheet, 2, 2)
	cell.Value = "Cell(2,2)"
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
