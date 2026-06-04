package sample

import (
	"strings"

	"github.com/xuri/excelize/v2"
)

func NewSingleSample(f *excelize.File, sheet string, receiptText ReceiptText) error {
	page, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}

	f.SetActiveSheet(page)
	// установка ширины ячеек
	if err := setCellsWidth(f, calcSingleWidth, sheet); err != nil {
		return err
	}
	if err := DrawSingleSampleAtOffset(f, sheet, receiptText, 0); err != nil {
		return err
	}
	// установка области печати
	if err := setPrintLayout(f, sheet); err != nil {
		return err
	}

	return nil
}

func NewDuoSample(f *excelize.File, sheet string, receiptText ReceiptText) error {
	page, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}

	f.SetActiveSheet(page)
	// установка ширины ячеек
	if err := setCellsWidth(f, CalcDuoWidth, sheet); err != nil {
		return err
	}
	if err := DrawDuoSampleAtOffset(f, sheet, receiptText, 0); err != nil {
		return err
	}
	// установка области печати
	if err := setPrintLayout(f, sheet); err != nil {
		return err
	}
	return nil

}

func NewPairSample(f *excelize.File, sheet string) error {
	page, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	f.SetActiveSheet(page)
	if err := setCellsWidth(f, CalcDuoWidth, sheet); err != nil {
		return err
	}
	return setPrintLayout(f, sheet)
}

func DrawSingleSampleAtOffset(f *excelize.File, sheet string, receiptText ReceiptText, rowOffset int) error {
	headers := headersWithReceiptText(CalcSingleHeaders, receiptText)

	// установка высоты ячеек
	if err := setCellsHeightAtOffset(f, calcHeight, sheet, rowOffset); err != nil {
		return err
	}
	// установка шрифта, размера, отрисовка границ
	if err := setSingleStyleAtOffset(f, sheet, rowOffset); err != nil {
		return err
	}
	// обьединение ячеек
	if err := mergeCellsAtOffset(f, headers, sheet, rowOffset); err != nil {
		return err
	}
	// установка хедеров и подвала
	if err := setHeadersAtOffset(f, headers, sheet, rowOffset); err != nil {
		return err
	}
	return nil
}

func DrawDuoSampleAtOffset(f *excelize.File, sheet string, receiptText ReceiptText, rowOffset int) error {
	headers := headersWithReceiptText(CalcDuoHeaders, receiptText)

	// установка высоты ячеек
	if err := setCellsHeightAtOffset(f, calcHeight, sheet, rowOffset); err != nil {
		return err
	}
	// установка шрифта, размераотрисовка границ
	if err := setDuoStyleAtOffset(f, sheet, rowOffset); err != nil {
		return err
	}
	// обьединение ячеек
	if err := mergeCellsAtOffset(f, headers, sheet, rowOffset); err != nil {
		return err
	}
	// установка хедеров и подвала
	if err := setHeadersAtOffset(f, headers, sheet, rowOffset); err != nil {
		return err
	}
	return nil
}

func setPrintLayout(f *excelize.File, sheet string) error {
	enable, zero := true, 0
	if err := f.SetSheetProps(sheet, &excelize.SheetPropsOptions{
		FitToPage: &enable,
	}); err != nil {
		return err
	}
	orientation := "landscape"
	if err := f.SetPageLayout(sheet, &excelize.PageLayoutOptions{
		Orientation: &orientation,
		FitToHeight: &zero,
	}); err != nil {
		return err
	}
	return nil
}

func setCellsWidth(f *excelize.File, w map[string]float64, sheet string) error {
	for i, v := range w {
		if err := f.SetColWidth(sheet, i, i, v); err != nil {
			return err
		}
	}
	return nil
}

func setCellsHeight(f *excelize.File, h map[int]float64, sheet string) error {
	return setCellsHeightAtOffset(f, h, sheet, 0)
}

func setCellsHeightAtOffset(f *excelize.File, h map[int]float64, sheet string, rowOffset int) error {
	for i, v := range h {
		if err := f.SetRowHeight(sheet, i+rowOffset, v); err != nil {
			return err
		}
	}
	return nil
}

func setSingleStyle(f *excelize.File, sheet string) error {
	return setSingleStyleAtOffset(f, sheet, 0)
}
func setSingleStyleAtOffset(f *excelize.File, sheet string, rowOffset int) error {
	if err := setMainStyle(f, shiftRange("B5:T7", rowOffset), sheet); err != nil {
		return err
	}
	if err := setLeftStyle(f, shiftRange("B7:D7", rowOffset), sheet, 1); err != nil {
		return err
	}
	// шапка
	if err := setLeftStyle(f, shiftRange("B1:U4", rowOffset), sheet, 0); err != nil {
		return err
	}
	// подвал
	if err := setLeftStyle(f, shiftRange("B9:B9", rowOffset), sheet, 0); err != nil {
		return err
	}
	if err := setValueStyle(f, shiftRange("E7:T7", rowOffset), sheet); err != nil {
		return err
	}
	return nil
}
func setMainStyle(f *excelize.File, cellRange, sheet string) error {
	style, err := f.NewStyle(baseStyle(7, "center", true, 1))
	if err != nil {
		return err
	}
	from, to := splitRange(cellRange)
	return f.SetCellStyle(sheet, from, to, style)
}
func setLeftStyle(f *excelize.File, cellRange, sheet string, border int) error {
	style, err := f.NewStyle(baseStyle(7, "left", true, border))
	if err != nil {
		return err
	}
	from, to := splitRange(cellRange)
	return f.SetCellStyle(sheet, from, to, style)
}
func setValueStyle(f *excelize.File, cellRange, sheet string) error {
	style, err := f.NewStyle(baseStyle(7, "right", true, 1))
	if err != nil {
		return err
	}
	from, to := splitRange(cellRange)
	return f.SetCellStyle(sheet, from, to, style)
}
func baseStyle(fontSize float64, horizontal string, wrap bool, border int) *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: border},
			{Type: "top", Color: "000000", Style: border},
			{Type: "bottom", Color: "000000", Style: border},
			{Type: "right", Color: "000000", Style: border},
			{Type: "horizontal", Color: "000000", Style: border},
			{Type: "vertical", Color: "000000", Style: border},
		},
		Font: &excelize.Font{
			Family: "Arial",
			Size:   fontSize,
		},
		Alignment: &excelize.Alignment{
			Horizontal: horizontal,
			Vertical:   "center",
			WrapText:   wrap,
		},
	}
}
func mergeCells(f *excelize.File, h map[string]string, sheet string) error {
	return mergeCellsAtOffset(f, h, sheet, 0)
}
func mergeCellsAtOffset(f *excelize.File, h map[string]string, sheet string, rowOffset int) error {
	for i := range h {
		if err := mergeCell(f, shiftRange(i, rowOffset), sheet); err != nil {
			return err
		}
	}
	return nil
}
func mergeCell(f *excelize.File, in, sheet string) error {
	parts := strings.Split(in, ":")
	if len(parts) == 2 {
		if err := f.MergeCell(sheet, parts[0], parts[1]); err != nil {
			return err
		}
	}
	return nil
}
func setHeaders(f *excelize.File, h map[string]string, sheet string) error {
	return setHeadersAtOffset(f, h, sheet, 0)
}
func setHeadersAtOffset(f *excelize.File, h map[string]string, sheet string, rowOffset int) error {
	for i, v := range h {
		parts := strings.Split(i, ":")
		cell, err := shiftCellRow(parts[0], rowOffset)
		if err != nil {
			return err
		}
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return err
		}
	}
	return nil
}
func setDuoStyle(f *excelize.File, sheet string) error {
	return setDuoStyleAtOffset(f, sheet, 0)
}
func setDuoStyleAtOffset(f *excelize.File, sheet string, rowOffset int) error {
	if err := setMainStyle(f, shiftRange("B5:U8", rowOffset), sheet); err != nil {
		return err
	}
	if err := setLeftStyle(f, shiftRange("B7:D8", rowOffset), sheet, 1); err != nil {
		return err
	}
	// шапка
	if err := setLeftStyle(f, shiftRange("B1:U4", rowOffset), sheet, 0); err != nil {
		return err
	}
	// подвал
	if err := setLeftStyle(f, shiftRange("B9:B9", rowOffset), sheet, 0); err != nil {
		return err
	}
	if err := setRatioStyle(f, shiftRange("H6:H6", rowOffset), sheet); err != nil {
		return err
	}
	if err := setValueStyle(f, shiftRange("E7:U8", rowOffset), sheet); err != nil {
		return err
	}
	if err := mergeCellsAtOffset(f, map[string]string{
		"B7:B8": "участок",
		"C7:C8": "имя",
		"R7:R8": "сумма",
		"S7:S8": "баланс",
		"T7:T8": "оплата",
		"U7:U8": "к оплате",
	}, sheet, rowOffset); err != nil {
		return err
	}
	return nil
}
func setRatioStyle(f *excelize.File, cellRange, sheet string) error {
	s := baseStyle(6, "center", true, 1)
	s.Alignment.TextRotation = 90
	style, err := f.NewStyle(s)
	if err != nil {
		return err
	}
	from, to := splitRange(cellRange)
	return f.SetCellStyle(sheet, from, to, style)
}

func shiftRange(cellRange string, rowOffset int) string {
	parts := strings.Split(cellRange, ":")
	if len(parts) == 1 {
		shifted, _ := shiftCellRow(parts[0], rowOffset)
		return shifted
	}

	from, _ := shiftCellRow(parts[0], rowOffset)
	to, _ := shiftCellRow(parts[1], rowOffset)
	return from + ":" + to
}

func shiftCellRow(cell string, rowOffset int) (string, error) {
	col, row, err := excelize.CellNameToCoordinates(cell)
	if err != nil {
		return "", err
	}
	return excelize.CoordinatesToCellName(col, row+rowOffset)
}

func splitRange(cellRange string) (string, string) {
	parts := strings.Split(cellRange, ":")
	if len(parts) == 1 {
		return parts[0], parts[0]
	}
	return parts[0], parts[1]
}
