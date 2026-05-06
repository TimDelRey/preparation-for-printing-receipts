package sample

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

const archiveSheet = "Архив"

type ReceiptText struct {
	HeaderLines [4]string
	Footer      string
}

func ReadReceiptText(f *excelize.File) (ReceiptText, error) {
	cells := []string{"B94", "B95", "B96", "B97"}
	var text ReceiptText

	for i, cell := range cells {
		value, err := f.GetCellValue(archiveSheet, cell)
		if err != nil {
			return ReceiptText{}, fmt.Errorf("read archive header %s: %w", cell, err)
		}
		text.HeaderLines[i] = value
	}

	footer, err := f.GetCellValue(archiveSheet, "B87")
	if err != nil {
		return ReceiptText{}, fmt.Errorf("read archive footer B87: %w", err)
	}
	text.Footer = footer

	return text, nil
}

func headersWithReceiptText(headers map[string]string, text ReceiptText) map[string]string {
	withText := make(map[string]string, len(headers)+5)
	for cell, value := range headers {
		withText[cell] = value
	}

	withText["B1:M1"] = text.HeaderLines[0]
	withText["B2:M2"] = text.HeaderLines[1]
	withText["B3:M3"] = text.HeaderLines[2]
	withText["B4:M4"] = text.HeaderLines[3]
	withText["B9:T9"] = text.Footer

	return withText
}
