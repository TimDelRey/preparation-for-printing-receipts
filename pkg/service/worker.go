package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"script_for_receipts/pkg/domain"
	"script_for_receipts/pkg/domain/sample"
	"time"
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
		// получение тарифа
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

	// удаление стартового листа
	if err := newFile.DeleteSheet("Sheet1"); err != nil {
		fmt.Println(err)
		return err
	}

	if err := newFile.SaveAs("Receipts.xlsx"); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func fragmentationReceips(file, newFile *excelize.File, tariffCell string) error {
	receipt := domain.NewReceipt(file, tariffCell)
	newSheet := fmt.Sprintf("%s-уч.%s", receipt.FullName, receipt.PlaceNumber)
	dateCell := "S1"
	dateVal, _ := file.GetCellValue(sheet, "T1")

	t, err := time.Parse("2006-01-02", dateVal)
	if err != nil {
		fmt.Println("не удалось распарсить дату:", dateVal)
		t = time.Now()
	}
	months := []string{
		"январь", "февраль", "март", "апрель", "май", "июнь",
		"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь",
	}
	formatted := fmt.Sprintf("%s.%d", months[int(t.Month())-1], t.Year())

	switch {
	case receipt.Single != nil:

		if err := sample.NewSingleSample(newFile, newSheet); err != nil {
			fmt.Println("не удалось создать шаблон")
		}
		if err := newFile.SetCellValue(newSheet, dateCell, formatted); err != nil {
			fmt.Println("не удалось установить дату")
		}
		if err := domain.PrintSingleReceipt(newFile, *receipt); err != nil {
			fmt.Println("не удалось распечатать стракт квитанции")
		}
		if err := file.RemoveRow(sheet, 10); err != nil {
			fmt.Println("не удалось  удалить квитанцию из рабочего документа")
		}
		return nil
	case receipt.Duo != nil:
		if err := sample.NewDuoSample(newFile, newSheet); err != nil {
			fmt.Println("не удалось создать шаблон")
		}
		if err := newFile.SetCellValue(newSheet, dateCell, formatted); err != nil {
			fmt.Println("не удалось установить дату")
		}
		if err := domain.PrintDuoReceipt(newFile, *receipt); err != nil {
			fmt.Println("не удалось распечатать стракт квитанции")
		}
		if err := file.RemoveRow(sheet, 10); err != nil {
			fmt.Println("не удалось  удалить квитанцию из рабочего документа")
		}
		if err := file.RemoveRow(sheet, 10); err != nil {
			fmt.Println("не удалось  удалить квитанцию из рабочего документа")
		}
		return nil
	default:
		return fmt.Errorf("some else error")
	}
}
