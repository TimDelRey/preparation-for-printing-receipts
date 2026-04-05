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
}
