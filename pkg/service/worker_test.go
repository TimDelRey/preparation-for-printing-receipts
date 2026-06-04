package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestRun_GeneratesReceiptsWorkbook(t *testing.T) {
	exePath, err := os.Executable()
	if err != nil {
		t.Fatalf("resolve test binary path: %v", err)
	}

	exeDir := filepath.Dir(exePath)
	inputPath := filepath.Join(exeDir, "work_doc_4.xlsx")
	outputPath := filepath.Join(exeDir, "Receipts.xlsx")

	sourceData, err := os.ReadFile(filepath.Join("..", "..", "bin", "work_doc_4.xlsx"))
	if err != nil {
		t.Fatalf("read fixture workbook: %v", err)
	}

	t.Cleanup(func() {
		_ = os.Remove(inputPath)
		_ = os.Remove(outputPath)
	})

	if err := os.WriteFile(inputPath, sourceData, 0o644); err != nil {
		t.Fatalf("prepare input workbook: %v", err)
	}

	if err := Run(); err != nil {
		t.Fatalf("run service: %v", err)
	}

	if _, err := os.Stat(outputPath); err != nil {
		t.Fatalf("receipts workbook was not created: %v", err)
	}

	resultFile, err := excelize.OpenFile(outputPath)
	if err != nil {
		t.Fatalf("open generated workbook: %v", err)
	}
	defer resultFile.Close()

	sheets := resultFile.GetSheetList()
	if len(sheets) == 0 {
		t.Fatal("generated workbook has no sheets")
	}

	if sheets[0] == "Sheet1" {
		t.Fatal("default sheet was not removed from generated workbook")
	}

	placeNumber, err := resultFile.GetCellValue(sheets[0], "B7")
	if err != nil {
		t.Fatalf("read generated receipt cell: %v", err)
	}
	if placeNumber == "" {
		t.Fatal("generated receipt does not contain place number in B7")
	}

	sourceFile, err := excelize.OpenFile(inputPath)
	if err != nil {
		t.Fatalf("open source workbook: %v", err)
	}
	defer sourceFile.Close()

	sourceReceiptCount, err := countReceipts(sourceFile)
	if err != nil {
		t.Fatalf("count source receipts: %v", err)
	}
	wantSheetCount := (sourceReceiptCount + 1) / 2
	if len(sheets) != wantSheetCount {
		t.Fatalf("generated workbook has %d sheets, want %d", len(sheets), wantSheetCount)
	}

	firstReceipt, err := readNextReceiptAtOffset(sourceFile, 0)
	if err != nil {
		t.Fatalf("read first source receipt: %v", err)
	}
	if firstReceipt == nil {
		t.Fatal("source workbook has no first receipt")
	}
	secondReceipt, err := readNextReceiptAtOffset(sourceFile, firstReceipt.rows)
	if err != nil {
		t.Fatalf("read second source receipt: %v", err)
	}
	if secondReceipt == nil {
		t.Fatal("source workbook has no second receipt")
	}

	wantFirstSheet := safeSheetName(receiptSheetName(firstReceipt.receipt, secondReceipt.receipt), map[string]int{})
	if sheets[0] != wantFirstSheet {
		t.Fatalf("first sheet name = %q, want %q", sheets[0], wantFirstSheet)
	}

	sourceHeader, err := sourceFile.GetCellValue("Архив", "B94")
	if err != nil {
		t.Fatalf("read source header: %v", err)
	}
	resultHeader, err := resultFile.GetCellValue(sheets[0], "B1")
	if err != nil {
		t.Fatalf("read generated header: %v", err)
	}
	if resultHeader != sourceHeader {
		t.Fatalf("generated header = %q, want %q", resultHeader, sourceHeader)
	}

	sourceFooter, err := sourceFile.GetCellValue("Архив", "B87")
	if err != nil {
		t.Fatalf("read source footer: %v", err)
	}
	resultFooter, err := resultFile.GetCellValue(sheets[0], "B9")
	if err != nil {
		t.Fatalf("read generated footer: %v", err)
	}
	if resultFooter != sourceFooter {
		t.Fatal("generated footer does not match source archive footer")
	}

	secondPlaceNumber, err := resultFile.GetCellValue(sheets[0], "B22")
	if err != nil {
		t.Fatalf("read second generated receipt cell: %v", err)
	}
	if secondPlaceNumber != secondReceipt.receipt.PlaceNumber {
		t.Fatalf("second generated place number = %q, want %q", secondPlaceNumber, secondReceipt.receipt.PlaceNumber)
	}

	secondHeader, err := resultFile.GetCellValue(sheets[0], "B16")
	if err != nil {
		t.Fatalf("read second generated header: %v", err)
	}
	if secondHeader != sourceHeader {
		t.Fatalf("second generated header = %q, want %q", secondHeader, sourceHeader)
	}

	secondFooter, err := resultFile.GetCellValue(sheets[0], "B24")
	if err != nil {
		t.Fatalf("read second generated footer: %v", err)
	}
	if secondFooter != sourceFooter {
		t.Fatal("second generated footer does not match source archive footer")
	}
}

func countReceipts(file *excelize.File) (int, error) {
	count := 0
	rowOffset := 0

	for {
		next, err := readNextReceiptAtOffset(file, rowOffset)
		if err != nil {
			return 0, err
		}
		if next == nil {
			return count, nil
		}

		count++
		rowOffset += next.rows
	}
}
