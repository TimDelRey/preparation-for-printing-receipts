package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"script_for_receipts/pkg/internal/domain"
	"script_for_receipts/pkg/internal/domain/sample"
)

func run() {
	file, err := excelize.OpenFile("work_doc_4.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	newFile := excelize.NewFile()
	defer newFile.Close()

	for {
		tariffCell, err := file.GetCellValue("Расчеты", "E10")
		if err != nil {
			fmt.Println(err)
			break
		}
		if tariffCell == "" {
			break
		}

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
	}

	switch receipt.Single.Tariff.TariffName {
	case "Тариф 1":
		if err := sample.NewSingleSample(newFile, receipt.FullName); err != nil {
			return err
		}
		if err := domain.PrintSingleReceipt(newFile, receipt); err != nil {
			return err
		}
		if err := file.RemoveRow("Расчеты", 10); err != nil {
			return err
		}
		return nil
	default:
		if err := sample.NewDuoSample(newFile, receipt.FullName); err != nil {
			return err
		}
		if err := domain.PrintDuoReceipt(newFile, receipt); err != nil {
			return err
		}
		if err := file.RemoveRow("Расчеты", 10); err != nil {
			return err
		}
		if err := file.RemoveRow("Расчеты", 11); err != nil {
			return err
		}
		return nil
	}
}
