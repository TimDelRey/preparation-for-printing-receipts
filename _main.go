package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	file, err := excelize.OpenFile("./doc4.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	cell, err := file.GetCellValue("Calc", "D31")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
}
