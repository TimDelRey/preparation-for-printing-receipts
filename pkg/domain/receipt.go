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

func PrintSingleReceipt(f *excelize.File, r Receipt) error {
	values := map[string]any{
		"B7": r.PlaceNumber,
		"C7": r.FullName,
		"D7": r.Single.Tariff.TariffName,
		"E7": r.Single.Tariff.CurrentInd,
		"F7": r.Single.Tariff.LastInd,
		"G7": r.Single.Tariff.DifValue,
		"H7": r.Single.Tariff.Step1Calc,
		"I7": r.Single.Tariff.Step2Calc,
		"J7": r.Single.Tariff.Step3Calc,
		"K7": r.Single.Tariff.Step1Price,
		"L7": r.Single.Tariff.Step2Price,
		"M7": r.Single.Tariff.Step3Price,
		"N7": r.Single.Tariff.Step1Arithmetic,
		"O7": r.Single.Tariff.Step2Arithmetic,
		"P7": r.Single.Tariff.Step3Arithmetic,
		"Q7": r.Summ,
		"R7": r.Balance,
		"S7": r.Payment,
		"T7": r.Debt,
	}
	newSheet := fmt.Sprintf("%s-уч.%s", r.FullName, r.PlaceNumber)

	for cell, val := range values {
		if err := f.SetCellValue(newSheet, cell, val); err != nil {
			return err
		}
	}

	if err := f.RemoveRow(newSheet, 8); err != nil {
		return err
	}
	return nil
}

func PrintDuoReceipt(f *excelize.File, r Receipt) error {
	values := map[string]any{
		"B7": r.PlaceNumber,
		"C7": r.FullName,
		"D7": r.Duo.T1.TariffName,
		"D8": r.Duo.T2.TariffName,
		"E7": r.Duo.T1.CurrentInd,
		"E8": r.Duo.T2.CurrentInd,
		"F7": r.Duo.T1.LastInd,
		"F8": r.Duo.T2.LastInd,
		"G7": r.Duo.T1.DifValue,
		"G8": r.Duo.T2.DifValue,
		"H7": r.Duo.T1.Ratio,
		"H8": r.Duo.T2.Ratio,
		"I7": r.Duo.T1.Step1Calc,
		"I8": r.Duo.T2.Step1Calc,
		"J7": r.Duo.T1.Step2Calc,
		"J8": r.Duo.T2.Step2Calc,
		"K7": r.Duo.T1.Step3Calc,
		"K8": r.Duo.T2.Step3Calc,
		"L7": r.Duo.T1.Step1Price,
		"L8": r.Duo.T2.Step1Price,
		"M7": r.Duo.T1.Step2Price,
		"M8": r.Duo.T2.Step2Price,
		"N7": r.Duo.T1.Step3Price,
		"N8": r.Duo.T2.Step3Price,
		"O7": r.Duo.T1.Step1Arithmetic,
		"O8": r.Duo.T2.Step1Arithmetic,
		"P7": r.Duo.T1.Step2Arithmetic,
		"P8": r.Duo.T2.Step2Arithmetic,
		"Q7": r.Duo.T1.Step3Arithmetic,
		"Q8": r.Duo.T2.Step3Arithmetic,
		"R7": r.Summ,
		"S7": r.Balance,
		"T7": r.Payment,
		"U7": r.Debt,
	}
	newSheet := fmt.Sprintf("%s-уч.%s", r.FullName, r.PlaceNumber)

	for cell, val := range values {
		if err := f.SetCellValue(newSheet, cell, val); err != nil {
			return err
		}
	}
	return nil
}
