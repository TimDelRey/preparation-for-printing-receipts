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
}
