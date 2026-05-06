package domain

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

const (
	tariffCell = "E10"
	sheet      = "Calc"
)

type Receipt struct {
	PlaceNumber string
	FullName    string

	Single *SingleZone
	Duo    *DuoCalc

	Summ    string
	Balance string
	Payment string
	Debt    string
}

type ReceiptSheetLayout struct {
	SingleFields  map[string]string
	DuoFields     map[string]string
	SingleTrimRow int
}

func DefaultReceiptSheetLayout() ReceiptSheetLayout {
	return ReceiptSheetLayout{
		SingleFields: map[string]string{
			"place_number":     "B7",
			"full_name":        "C7",
			"tariff_name":      "D7",
			"current_ind":      "E7",
			"last_ind":         "F7",
			"dif_value":        "G7",
			"step1_calc":       "H7",
			"step2_calc":       "I7",
			"step3_calc":       "J7",
			"step1_price":      "K7",
			"step2_price":      "L7",
			"step3_price":      "M7",
			"step1_arithmetic": "N7",
			"step2_arithmetic": "O7",
			"step3_arithmetic": "P7",
			"summ":             "Q7",
			"balance":          "R7",
			"payment":          "S7",
			"debt":             "T7",
		},
		DuoFields: map[string]string{
			"place_number":        "B7",
			"full_name":           "C7",
			"t1_tariff_name":      "D7",
			"t2_tariff_name":      "D8",
			"t1_current_ind":      "E7",
			"t2_current_ind":      "E8",
			"t1_last_ind":         "F7",
			"t2_last_ind":         "F8",
			"t1_dif_value":        "G7",
			"t2_dif_value":        "G8",
			"t1_ratio":            "H7",
			"t2_ratio":            "H8",
			"t1_step1_calc":       "I7",
			"t2_step1_calc":       "I8",
			"t1_step2_calc":       "J7",
			"t2_step2_calc":       "J8",
			"t1_step3_calc":       "K7",
			"t2_step3_calc":       "K8",
			"t1_step1_price":      "L7",
			"t2_step1_price":      "L8",
			"t1_step2_price":      "M7",
			"t2_step2_price":      "M8",
			"t1_step3_price":      "N7",
			"t2_step3_price":      "N8",
			"t1_step1_arithmetic": "O7",
			"t2_step1_arithmetic": "O8",
			"t1_step2_arithmetic": "P7",
			"t2_step2_arithmetic": "P8",
			"t1_step3_arithmetic": "Q7",
			"t2_step3_arithmetic": "Q8",
			"summ":                "R7",
			"balance":             "S7",
			"payment":             "T7",
			"debt":                "U7",
		},
		SingleTrimRow: 8,
	}
}

type SingleZone struct {
	Tariff TariffCalc
}

type DuoCalc struct {
	T1 TariffCalc
	T2 TariffCalc
}

type TariffCalc struct {
	TariffName      string
	CurrentInd      string
	LastInd         string
	DifValue        string
	Ratio           string
	Step1Calc       string
	Step2Calc       string
	Step3Calc       string
	Step1Price      string
	Step2Price      string
	Step3Price      string
	Step1Arithmetic string
	Step2Arithmetic string
	Step3Arithmetic string
}

func NewReceipt(file *excelize.File, tariff string) *Receipt {
	p := &Receipt{}

	p.PlaceNumber, _ = file.GetCellValue(sheet, cellRel(tariffCell, -2, 0))
	p.FullName, _ = file.GetCellValue(sheet, cellRel(tariffCell, -1, 0))
	p.Summ, _ = file.GetCellValue(sheet, cellRel(tariffCell, 14, 0))
	p.Balance, _ = file.GetCellValue(sheet, cellRel(tariffCell, 15, 0))
	p.Payment, _ = file.GetCellValue(sheet, cellRel(tariffCell, 16, 0))
	p.Debt, _ = file.GetCellValue(sheet, cellRel(tariffCell, 17, 0))

	tariffName, _ := file.GetCellValue(sheet, tariffCell)
	currentInd, _ := file.GetCellValue(sheet, cellRel(tariffCell, 1, 0))
	lastInd, _ := file.GetCellValue(sheet, cellRel(tariffCell, 2, 0))
	difValue, _ := file.GetCellValue(sheet, cellRel(tariffCell, 3, 0))

	step1Calc, _ := file.GetCellValue(sheet, cellRel(tariffCell, 5, 0))
	step2Calc, _ := file.GetCellValue(sheet, cellRel(tariffCell, 6, 0))
	step3Calc, _ := file.GetCellValue(sheet, cellRel(tariffCell, 7, 0))

	step1Price, _ := file.GetCellValue(sheet, cellRel(tariffCell, 8, 0))
	step2Price, _ := file.GetCellValue(sheet, cellRel(tariffCell, 9, 0))
	step3Price, _ := file.GetCellValue(sheet, cellRel(tariffCell, 10, 0))

	step1Arithmetic, _ := file.GetCellValue(sheet, cellRel(tariffCell, 11, 0))
	step2Arithmetic, _ := file.GetCellValue(sheet, cellRel(tariffCell, 12, 0))
	step3Arithmetic, _ := file.GetCellValue(sheet, cellRel(tariffCell, 13, 0))

	switch tariff {
	case "Тариф 1":
		p.Single = &SingleZone{
			Tariff: TariffCalc{
				TariffName:      tariffName,
				CurrentInd:      currentInd,
				LastInd:         lastInd,
				DifValue:        difValue,
				Step1Calc:       step1Calc,
				Step2Calc:       step2Calc,
				Step3Calc:       step3Calc,
				Step1Price:      step1Price,
				Step2Price:      step2Price,
				Step3Price:      step3Price,
				Step1Arithmetic: step1Arithmetic,
				Step2Arithmetic: step2Arithmetic,
				Step3Arithmetic: step3Arithmetic,
			},
		}

	default:
		ratioT1, _ := file.GetCellValue(sheet, cellRel(tariffCell, 4, 0))
		ratioT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 4, 1))

		tariffNameT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 0, 1))
		currentIndT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 1, 1))
		lastIndT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 2, 1))
		difValueT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 3, 1))

		step1CalcT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 5, 1))
		step2CalcT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 6, 1))
		step3CalcT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 7, 1))

		step1PriceT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 8, 1))
		step2PriceT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 9, 1))
		step3PriceT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 10, 1))

		step1ArithmeticT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 11, 1))
		step2ArithmeticT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 12, 1))
		step3ArithmeticT2, _ := file.GetCellValue(sheet, cellRel(tariffCell, 13, 1))

		p.Duo = &DuoCalc{
			T1: TariffCalc{
				TariffName:      tariffName,
				CurrentInd:      currentInd,
				LastInd:         lastInd,
				DifValue:        difValue,
				Ratio:           ratioT1,
				Step1Calc:       step1Calc,
				Step2Calc:       step2Calc,
				Step3Calc:       step3Calc,
				Step1Price:      step1Price,
				Step2Price:      step2Price,
				Step3Price:      step3Price,
				Step1Arithmetic: step1Arithmetic,
				Step2Arithmetic: step2Arithmetic,
				Step3Arithmetic: step3Arithmetic,
			},
			T2: TariffCalc{
				TariffName:      tariffNameT2,
				CurrentInd:      currentIndT2,
				LastInd:         lastIndT2,
				DifValue:        difValueT2,
				Ratio:           ratioT2,
				Step1Calc:       step1CalcT2,
				Step2Calc:       step2CalcT2,
				Step3Calc:       step3CalcT2,
				Step1Price:      step1PriceT2,
				Step2Price:      step2PriceT2,
				Step3Price:      step3PriceT2,
				Step1Arithmetic: step1ArithmeticT2,
				Step2Arithmetic: step2ArithmeticT2,
				Step3Arithmetic: step3ArithmeticT2,
			},
		}
	}

	return p
}

func cellRel(mainCell string, difCol, difRow int) string {
	col, row, _ := excelize.CellNameToCoordinates(mainCell)
	newCell, _ := excelize.CoordinatesToCellName(col+difCol, row+difRow)
	return newCell
}

func (r Receipt) SheetName() string {
	return fmt.Sprintf("%s-уч.%s", r.FullName, r.PlaceNumber)
}

func PrintSingleReceipt(f *excelize.File, r Receipt) error {
	return PrintSingleReceiptAtOffset(f, r, DefaultReceiptSheetLayout(), 0)
}

func PrintSingleReceiptAtOffset(f *excelize.File, r Receipt, layout ReceiptSheetLayout, rowOffset int) error {
	if err := fillMappedValues(f, r.SheetName(), layout.SingleFields, singleReceiptValues(r), rowOffset); err != nil {
		return err
	}

	if err := f.RemoveRow(r.SheetName(), layout.SingleTrimRow+rowOffset); err != nil {
		return err
	}
	return nil
}

func PrintDuoReceipt(f *excelize.File, r Receipt) error {
	return PrintDuoReceiptAtOffset(f, r, DefaultReceiptSheetLayout(), 0)
}

func PrintDuoReceiptAtOffset(f *excelize.File, r Receipt, layout ReceiptSheetLayout, rowOffset int) error {
	return fillMappedValues(f, r.SheetName(), layout.DuoFields, duoReceiptValues(r), rowOffset)
}

func singleReceiptValues(r Receipt) map[string]any {
	return map[string]any{
		"place_number":     r.PlaceNumber,
		"full_name":        r.FullName,
		"tariff_name":      r.Single.Tariff.TariffName,
		"current_ind":      r.Single.Tariff.CurrentInd,
		"last_ind":         r.Single.Tariff.LastInd,
		"dif_value":        r.Single.Tariff.DifValue,
		"step1_calc":       r.Single.Tariff.Step1Calc,
		"step2_calc":       r.Single.Tariff.Step2Calc,
		"step3_calc":       r.Single.Tariff.Step3Calc,
		"step1_price":      r.Single.Tariff.Step1Price,
		"step2_price":      r.Single.Tariff.Step2Price,
		"step3_price":      r.Single.Tariff.Step3Price,
		"step1_arithmetic": r.Single.Tariff.Step1Arithmetic,
		"step2_arithmetic": r.Single.Tariff.Step2Arithmetic,
		"step3_arithmetic": r.Single.Tariff.Step3Arithmetic,
		"summ":             r.Summ,
		"balance":          r.Balance,
		"payment":          r.Payment,
		"debt":             r.Debt,
	}
}

func duoReceiptValues(r Receipt) map[string]any {
	return map[string]any{
		"place_number":        r.PlaceNumber,
		"full_name":           r.FullName,
		"t1_tariff_name":      r.Duo.T1.TariffName,
		"t2_tariff_name":      r.Duo.T2.TariffName,
		"t1_current_ind":      r.Duo.T1.CurrentInd,
		"t2_current_ind":      r.Duo.T2.CurrentInd,
		"t1_last_ind":         r.Duo.T1.LastInd,
		"t2_last_ind":         r.Duo.T2.LastInd,
		"t1_dif_value":        r.Duo.T1.DifValue,
		"t2_dif_value":        r.Duo.T2.DifValue,
		"t1_ratio":            r.Duo.T1.Ratio,
		"t2_ratio":            r.Duo.T2.Ratio,
		"t1_step1_calc":       r.Duo.T1.Step1Calc,
		"t2_step1_calc":       r.Duo.T2.Step1Calc,
		"t1_step2_calc":       r.Duo.T1.Step2Calc,
		"t2_step2_calc":       r.Duo.T2.Step2Calc,
		"t1_step3_calc":       r.Duo.T1.Step3Calc,
		"t2_step3_calc":       r.Duo.T2.Step3Calc,
		"t1_step1_price":      r.Duo.T1.Step1Price,
		"t2_step1_price":      r.Duo.T2.Step1Price,
		"t1_step2_price":      r.Duo.T1.Step2Price,
		"t2_step2_price":      r.Duo.T2.Step2Price,
		"t1_step3_price":      r.Duo.T1.Step3Price,
		"t2_step3_price":      r.Duo.T2.Step3Price,
		"t1_step1_arithmetic": r.Duo.T1.Step1Arithmetic,
		"t2_step1_arithmetic": r.Duo.T2.Step1Arithmetic,
		"t1_step2_arithmetic": r.Duo.T1.Step2Arithmetic,
		"t2_step2_arithmetic": r.Duo.T2.Step2Arithmetic,
		"t1_step3_arithmetic": r.Duo.T1.Step3Arithmetic,
		"t2_step3_arithmetic": r.Duo.T2.Step3Arithmetic,
		"summ":                r.Summ,
		"balance":             r.Balance,
		"payment":             r.Payment,
		"debt":                r.Debt,
	}
}

func fillMappedValues(f *excelize.File, sheet string, cells map[string]string, values map[string]any, rowOffset int) error {
	for key, cell := range cells {
		value, ok := values[key]
		if !ok {
			continue
		}

		shiftedCell, err := shiftCellRow(cell, rowOffset)
		if err != nil {
			return err
		}

		if err := f.SetCellValue(sheet, shiftedCell, value); err != nil {
			return err
		}
	}

	return nil
}

func shiftCellRow(cell string, rowOffset int) (string, error) {
	col, row, err := excelize.CellNameToCoordinates(cell)
	if err != nil {
		return "", err
	}

	return excelize.CoordinatesToCellName(col, row+rowOffset)
}
