package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
	"script_for_receipts/pkg/domain"
	"script_for_receipts/pkg/domain/sample"
	"time"
)

type Config struct {
	InputFileName  string
	OutputFileName string
	InputSheetName string
	TariffCell     string
	SourceDateCell string
	OutputDateCell string
	ReceiptLayout  domain.ReceiptLayout
}

func DefaultConfig() Config {
	return Config{
		InputFileName:  "work_doc_4.xlsx",
		OutputFileName: "Receipts.xlsx",
		InputSheetName: "Calc",
		TariffCell:     "E10",
		SourceDateCell: "T1",
		OutputDateCell: "S1",
		ReceiptLayout:  domain.DefaultReceiptLayout(),
	}
}

func Run() error {
	return RunWithConfig(DefaultConfig())
}

func RunWithConfig(cfg Config) error {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	xlsxPath := filepath.Join(exeDir, cfg.InputFileName)

	file, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	newFile := excelize.NewFile()
	defer newFile.Close()

	for {
		tariffCell, err := file.GetCellValue(cfg.InputSheetName, cfg.TariffCell)
		if err != nil {
			fmt.Println(err)
			break
		}
		if tariffCell == "" {
			break
		}

		if err := fragmentationReceips(file, newFile, cfg, tariffCell); err != nil {
			fmt.Println(err)
			return err
		}
	}

	// удаление стартового листа
	if err := newFile.DeleteSheet("Sheet1"); err != nil {
		fmt.Println(err)
		return err
	}

	outPath := filepath.Join(exeDir, cfg.OutputFileName)
	if err := newFile.SaveAs(outPath); err != nil {
		fmt.Println("Failed to save Excel:", err)
		return err
	}
	return nil
}

func fragmentationReceips(file, newFile *excelize.File, cfg Config, tariffCell string) error {
	receipt := domain.NewReceipt(file, cfg.ReceiptLayout, tariffCell)
	newSheet := receipt.SheetName()
	formatted := receiptDate(file, cfg)

	switch {
	case receipt.Single != nil:
		if err := createSingleReceipt(newFile, newSheet, formatted, cfg, *receipt); err != nil {
			return err
		}
		if err := file.RemoveRow(cfg.InputSheetName, 10); err != nil {
			fmt.Println("не удалось  удалить квитанцию из рабочего документа")
		}
		return nil
	case receipt.Duo != nil:
		if err := createDuoReceipt(newFile, newSheet, formatted, cfg, *receipt); err != nil {
			return err
		}
		if err := file.RemoveRow(cfg.InputSheetName, 10); err != nil {
			fmt.Println("не удалось  удалить квитанцию из рабочего документа")
		}
		if err := file.RemoveRow(cfg.InputSheetName, 10); err != nil {
			fmt.Println("не удалось  удалить квитанцию из рабочего документа")
		}
		return nil
	default:
		return fmt.Errorf("some else error")
	}
}

func createSingleReceipt(newFile *excelize.File, sheetName, formattedDate string, cfg Config, receipt domain.Receipt) error {
	if err := sample.NewSingleSample(newFile, sheetName); err != nil {
		return fmt.Errorf("не удалось создать шаблон")
	}
	if err := newFile.SetCellValue(sheetName, cfg.OutputDateCell, formattedDate); err != nil {
		return fmt.Errorf("не удалось установить дату")
	}
	if err := domain.PrintSingleReceipt(newFile, receipt); err != nil {
		return fmt.Errorf("не удалось распечатать стракт квитанции")
	}
	return nil
}

func createDuoReceipt(newFile *excelize.File, sheetName, formattedDate string, cfg Config, receipt domain.Receipt) error {
	if err := sample.NewDuoSample(newFile, sheetName); err != nil {
		return fmt.Errorf("не удалось создать шаблон")
	}
	if err := newFile.SetCellValue(sheetName, cfg.OutputDateCell, formattedDate); err != nil {
		return fmt.Errorf("не удалось установить дату")
	}
	if err := domain.PrintDuoReceipt(newFile, receipt); err != nil {
		return fmt.Errorf("не удалось распечатать стракт квитанции")
	}
	return nil
}

func receiptDate(file *excelize.File, cfg Config) string {
	dateVal, _ := file.GetCellValue(cfg.InputSheetName, cfg.SourceDateCell)

	t, err := time.Parse("2006-01-02", dateVal)
	if err != nil {
		fmt.Println("не удалось распарсить дату:", dateVal)
		t = time.Now()
	}

	months := []string{
		"январь", "февраль", "март", "апрель", "май", "июнь",
		"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь",
	}

	return fmt.Sprintf("%s.%d", months[int(t.Month())-1], t.Year())
}
