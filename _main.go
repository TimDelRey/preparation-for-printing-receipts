package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func cellRel(mainCell string, difCol, difRow int) string {
	col, row, _ := excelize.CellNameToCoordinates(mainCell)
	newCell, _ := excelize.CoordinatesToCellName(col+difCol, row+difRow)
	return newCell
}

func main() {
	file, err := excelize.OpenFile("./work_doc_4.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cell, err := file.GetCellValue("Calc", cellRel("E10", 0, 1))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
}
