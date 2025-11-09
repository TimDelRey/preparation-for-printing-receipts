package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"script_for_receipts/pkg/domain"
	"script_for_receipts/pkg/domain/sample"
)

const sheet = "Calc"

func Run() error {
	file, err := excelize.OpenFile("work_doc_4.xlsx")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	newFile := excelize.NewFile()
	defer newFile.Close()

	for {
		tariffCell, err := file.GetCellValue(sheet, "E10")
		if err != nil {
			fmt.Println(err)
			break
		}
		if tariffCell == "" {
			break
		}

		if err := fragmentationReceips(file, newFile, tariffCell); err != nil {
			fmt.Println(err)
			return err
		}
	}

	if err := newFile.SaveAs("Receipts.xlsx"); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func fragmentationReceips(file, newFile *excelize.File, tariffCell string) error {
	receipt := domain.NewReceipt(file, tariffCell)
	index, err := newFile.NewSheet(receipt.FullName)
	if err != nil {
		fmt.Println(err)
	}
	file.SetActiveSheet(index)

	switch {
	case receipt.Single != nil:
		if err := sample.NewSingleSample(newFile, receipt.FullName); err != nil {
			return err
		}
		if err := domain.PrintSingleReceipt(newFile, *receipt); err != nil {
			return err
		}
		if err := file.RemoveRow(sheet, 10); err != nil {
			return err
		}
		return nil
	case receipt.Duo != nil:
		if err := sample.NewDuoSample(newFile, receipt.FullName); err != nil {
			return err
		}
		if err := domain.PrintDuoReceipt(newFile, *receipt); err != nil {
			return err
		}
		if err := file.RemoveRow(sheet, 10); err != nil {
			return err
		}
		if err := file.RemoveRow(sheet, 11); err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("some else error")
	}
}
