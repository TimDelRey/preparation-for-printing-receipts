package sample

import (
	"github.com/xuri/excelize/v2"
	"strings"
)

func NewSingleSample(f *excelize.File, sheet string) error {
	page, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}

	f.SetActiveSheet(page)
	// установка ширины ячеек
	if err := setCellsWidth(f, calcSingleWidth, sheet); err != nil {
		return err
	}
	// установка высоты ячеек
	if err := setCellsHeight(f, calcHeight, sheet); err != nil {
		return err
	}
	// установка шрифта, размера, отрисовка границ
	if err := setSingleStyle(f, sheet); err != nil {
		return err
	}
	// обьединение ячеек
	if err := mergeCells(f, CalcSingleHeaders, sheet); err != nil {
		return err
	}
	// установка хедеров и подвала
	if err := setHeaders(f, CalcSingleHeaders, sheet); err != nil {
		return err
	}
	return nil
}

func NewDuoSample(f *excelize.File, sheet string) error {
	page, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}

	f.SetActiveSheet(page)
	// установка ширины ячеек
	if err := setCellsWidth(f, CalcDuoWidth, sheet); err != nil {
		return err
	}
	// установка высоты ячеек
	if err := setCellsHeight(f, calcHeight, sheet); err != nil {
		return err
	}
	// установка шрифта, размераотрисовка границ
	if err := setDuoStyle(f, sheet); err != nil {
		return err
	}
	// обьединение ячеек
	if err := mergeCells(f, CalcDuoHeaders, sheet); err != nil {
		return err
	}
	// установка хедеров и подвала
	if err := setHeaders(f, CalcDuoHeaders, sheet); err != nil {
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
	for i, v := range h {
		if err := f.SetRowHeight(sheet, i, v); err != nil {
			return err
		}
	}
	return nil
}

func setSingleStyle(f *excelize.File, sheet string) error {
	if err := setMainStyle(f, "B5", "T7", sheet); err != nil {
		return err
	}
	if err := setLeftStyle(f, "B7", "D7", sheet, 1); err != nil {
		return err
	}
	// шапка
	if err := setLeftStyle(f, "B1", "U4", sheet, 0); err != nil {
		return err
	}
	// подвал
	if err := setLeftStyle(f, "B9", "B9", sheet, 0); err != nil {
		return err
	}
	if err := setValueStyle(f, "E7", "T7", sheet); err != nil {
		return err
	}
	return nil
}
func setMainStyle(f *excelize.File, from, to, sheet string) error {
	style, err := f.NewStyle(baseStyle(7, "center", true, 1))
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, from, to, style)
}
func setLeftStyle(f *excelize.File, from, to, sheet string, border int) error {
	style, err := f.NewStyle(baseStyle(7, "left", true, border))
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, from, to, style)
}
func setValueStyle(f *excelize.File, from, to, sheet string) error {
	style, err := f.NewStyle(baseStyle(7, "right", true, 1))
	if err != nil {
		return err
	}
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
	for i := range h {
		if err := mergeCell(f, i, sheet); err != nil {
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
	for i, v := range h {
		parts := strings.Split(i, ":")
		if err := f.SetCellValue(sheet, parts[0], v); err != nil {
			return err
		}
	}
	return nil
}
func setDuoStyle(f *excelize.File, sheet string) error {
	if err := setMainStyle(f, "B5", "U8", sheet); err != nil {
		return err
	}
	if err := setLeftStyle(f, "B7", "D8", sheet, 1); err != nil {
		return err
	}
	// шапка
	if err := setLeftStyle(f, "B1", "U4", sheet, 0); err != nil {
		return err
	}
	// подвал
	if err := setLeftStyle(f, "B9", "B9", sheet, 0); err != nil {
		return err
	}
	if err := setRatioStyle(f, "H6", "H6", sheet); err != nil {
		return err
	}
	if err := setValueStyle(f, "E7", "U8", sheet); err != nil {
		return err
	}
	if err := mergeCells(f, map[string]string{
		"B7:B8": "участок",
		"C7:C8": "имя",
		"R7:R8": "сумма",
		"S7:S8": "баланс",
		"T7:T8": "оплата",
		"U7:U8": "к оплате",
	}, sheet); err != nil {
		return err
	}
	return nil
}
func setRatioStyle(f *excelize.File, from, to, sheet string) error {
	s := baseStyle(6, "center", true, 1)
	s.Alignment.TextRotation = 90
	style, err := f.NewStyle(s)
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, from, to, style)
}
