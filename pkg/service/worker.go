package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"

	"script_for_receipts/pkg/domain"
	"script_for_receipts/pkg/domain/sample"
)

const sheet = "Calc"
const firstTariffCell = "E10"
const dateCell = "S1"
const secondReceiptRowOffset = 15
const excelSheetNameMaxLength = 31

type nextReceipt struct {
	receipt *domain.Receipt
	rows    int
}

func Run() error {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	xlsxPath := filepath.Join(exeDir, "work_doc_4.xlsx")

	file, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	newFile := excelize.NewFile()
	defer newFile.Close()

	receiptText, err := sample.ReadReceiptText(file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	usedSheetNames := make(map[string]int)
	for {
		first, err := readNextReceiptAtOffset(file, 0)
		if err != nil {
			return err
		}
		if first == nil {
			break
		}

		second, err := readNextReceiptAtOffset(file, first.rows)
		if err != nil {
			return err
		}

		date, err := formattedReceiptDate(file)
		if err != nil {
			return err
		}

		sheetName := receiptSheetName(first.receipt, nil)
		if second != nil {
			sheetName = receiptSheetName(first.receipt, second.receipt)
		}
		sheetName = safeSheetName(sheetName, usedSheetNames)

		if err := printReceiptPair(newFile, sheetName, receiptText, date, first.receipt, optionalReceipt(second)); err != nil {
			return err
		}

		rowsToRemove := first.rows
		if second != nil {
			rowsToRemove += second.rows
		}
		if err := removeProcessedRows(file, rowsToRemove); err != nil {
			return err
		}
	}

	// удаление стартового листа
	if err := newFile.DeleteSheet("Sheet1"); err != nil {
		fmt.Println(err)
		return err
	}

	outPath := filepath.Join(exeDir, "Receipts.xlsx")
	if err := newFile.SaveAs(outPath); err != nil {
		fmt.Println("Failed to save Excel:", err)
		return err
	}
	return nil
}

func readNextReceipt(file *excelize.File) (*domain.Receipt, int, error) {
	next, err := readNextReceiptAtOffset(file, 0)
	if err != nil || next == nil {
		return nil, 0, err
	}
	return next.receipt, next.rows, nil
}

func readNextReceiptAtOffset(file *excelize.File, rowOffset int) (*nextReceipt, error) {
	tariffCell, err := shiftCellRow(firstTariffCell, rowOffset)
	if err != nil {
		return nil, err
	}

	tariff, err := file.GetCellValue(sheet, tariffCell)
	if err != nil {
		return nil, err
	}
	if tariff == "" {
		return nil, nil
	}

	receipt := domain.NewReceiptAtCell(file, tariffCell, tariff)
	rows := 2
	if receipt.Single != nil {
		rows = 1
	}
	return &nextReceipt{
		receipt: receipt,
		rows:    rows,
	}, nil
}

func formattedReceiptDate(file *excelize.File) (string, error) {
	dateVal, err := file.GetCellValue(sheet, "T1")
	if err != nil {
		return "", err
	}

	t, err := time.Parse("2006-01-02", dateVal)
	if err != nil {
		fmt.Println("не удалось распарсить дату:", dateVal)
		t = time.Now()
	}
	months := []string{
		"январь", "февраль", "март", "апрель", "май", "июнь",
		"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь",
	}
	return fmt.Sprintf("%s.%d", months[int(t.Month())-1], t.Year()), nil
}

func printReceiptPair(newFile *excelize.File, sheetName string, receiptText sample.ReceiptText, date string, first, second *domain.Receipt) error {
	if err := sample.NewPairSample(newFile, sheetName); err != nil {
		return err
	}

	if err := drawReceipt(newFile, sheetName, receiptText, date, first, 0); err != nil {
		return err
	}

	if second != nil {
		if err := drawReceipt(newFile, sheetName, receiptText, date, second, secondReceiptRowOffset); err != nil {
			return err
		}
	}

	return nil
}

func drawReceipt(newFile *excelize.File, sheetName string, receiptText sample.ReceiptText, date string, receipt *domain.Receipt, rowOffset int) error {
	if receipt.Single != nil {
		if err := sample.DrawSingleSampleAtOffset(newFile, sheetName, receiptText, rowOffset); err != nil {
			return err
		}
		if err := setReceiptDate(newFile, sheetName, date, rowOffset); err != nil {
			return err
		}
		return domain.PrintSingleReceiptOnSheet(newFile, sheetName, *receipt, rowOffset)
	}

	if receipt.Duo != nil {
		if err := sample.DrawDuoSampleAtOffset(newFile, sheetName, receiptText, rowOffset); err != nil {
			return err
		}
		if err := setReceiptDate(newFile, sheetName, date, rowOffset); err != nil {
			return err
		}
		return domain.PrintDuoReceiptOnSheet(newFile, sheetName, *receipt, rowOffset)
	}

	return fmt.Errorf("unknown receipt type")
}

func setReceiptDate(file *excelize.File, sheetName, date string, rowOffset int) error {
	cell, err := shiftCellRow(dateCell, rowOffset)
	if err != nil {
		return err
	}
	return file.SetCellValue(sheetName, cell, date)
}

func removeProcessedRows(file *excelize.File, rows int) error {
	for i := 0; i < rows; i++ {
		if err := file.RemoveRow(sheet, 10); err != nil {
			return err
		}
	}
	return nil
}

func optionalReceipt(next *nextReceipt) *domain.Receipt {
	if next == nil {
		return nil
	}
	return next.receipt
}

func receiptSheetName(first, second *domain.Receipt) string {
	name := receiptPlaceName(first)
	if second != nil {
		name += "-" + receiptPlaceName(second)
	}
	return name
}

func receiptPlaceName(receipt *domain.Receipt) string {
	return fmt.Sprintf("уч.%s", receipt.PlaceNumber)
}

func safeSheetName(name string, used map[string]int) string {
	replacer := strings.NewReplacer(
		":", "-",
		"\\", "-",
		"/", "-",
		"?", "-",
		"*", "-",
		"[", "-",
		"]", "-",
	)
	base := strings.TrimSpace(replacer.Replace(name))
	if base == "" {
		base = "Receipts"
	}
	base = trimSheetName(base, excelSheetNameMaxLength)

	count := used[base]
	used[base] = count + 1
	if count == 0 {
		return base
	}

	suffix := fmt.Sprintf("-%d", count+1)
	return trimSheetName(base, excelSheetNameMaxLength-len([]rune(suffix))) + suffix
}

func trimSheetName(name string, maxLength int) string {
	runes := []rune(name)
	if len(runes) <= maxLength {
		return name
	}
	return string(runes[:maxLength])
}

func shiftCellRow(cell string, rowOffset int) (string, error) {
	col, row, err := excelize.CellNameToCoordinates(cell)
	if err != nil {
		return "", err
	}
	return excelize.CoordinatesToCellName(col, row+rowOffset)
}
