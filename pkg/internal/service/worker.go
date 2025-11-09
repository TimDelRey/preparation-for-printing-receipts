package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"script_for_receipts/pkg/internal/domain"
	"script_for_receipts/pkg/internal/domain/sample"
)

func run() {
	// читаем файл
	file, err := excelize.OpenFile("./work_doc_4.xlsx")
	// file, err := excelize.OpenFile("work_doc_4.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// создаем новый
	newFile := excelize.NewFile()
	defer func() {
		if err := newFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	tariffCell, err := file.GetCellValue("Расчеты", "E10")
	if err != nil {
		fmt.Println(err)
		return
	}
	for tariffCell != nil {
		if err := fragmentationReceips(file, newFile, tariffCell); err != nil {
			fmt.Println(err)
		}
	}
	if err := newFile.SaveAs("Receipts.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func fragmentationReceips(file, newFile *excelize.File, tariffCell string) error {
	receipt := domain.NewReceipt(file, tariffCell)
	if _, err := newFile.NewSheet(receipt.FullName); err != nil {
		fmt.Println(err)
		return
	}

	switch receipt.TariffName {
	case "Тариф 1":
		err := sample.newSingleSample(newFile, receipt.FullName)
		if err != nil {
			fmt.Println(err)
			return
		}
		err := domain.printSingleReceipt(newFile, receipt)
		if err != nil {
			fmt.Println(err)
			return
		}
		err := file.RemoveRow("Расчеты", 10)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		err := sample.newDuoSample(newFile, receipt.FullName)
		if err != nil {
			fmt.Println(err)
			return
		}
		err := domain.printDuoReceipt(newFile, receipt)
		if err != nil {
			fmt.Println(err)
			return
		}
		err := file.RemoveRow("Расчеты", 10)
		if err != nil {
			fmt.Println(err)
			return
		}
		err := file.RemoveRow("Расчеты", 11)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
