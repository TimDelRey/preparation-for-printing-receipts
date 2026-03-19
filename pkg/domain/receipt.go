package domain

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type CellReader interface {
	GetCellValue(sheet, cell string, opts ...excelize.Options) (string, error)
}

type CellWriter interface {
	SetCellValue(sheet, cell string, value any) error
	RemoveRow(sheet string, row int) error
}

type ReceiptLayout struct {
	SheetName        string
	BaseTariffCell   string
	SingleTariffName string
}

func DefaultReceiptLayout() ReceiptLayout {
	return ReceiptLayout{
		SheetName:        "Calc",
		BaseTariffCell:   "E10",
		SingleTariffName: "Тариф 1",
	}
}

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

func NewReceipt(file CellReader, layout ReceiptLayout, tariff string) *Receipt {
	p := &Receipt{
		PlaceNumber: getCellValue(file, layout, -2, 0),
		FullName:    getCellValue(file, layout, -1, 0),
		Summ:        getCellValue(file, layout, 14, 0),
		Balance:     getCellValue(file, layout, 15, 0),
		Payment:     getCellValue(file, layout, 16, 0),
		Debt:        getCellValue(file, layout, 17, 0),
	}

	switch tariff {
	case layout.SingleTariffName:
		p.Single = &SingleZone{
			Tariff: newTariffCalc(file, layout, 0),
		}
	default:
		p.Duo = &DuoCalc{
			T1: newTariffCalc(file, layout, 0),
			T2: newTariffCalc(file, layout, 1),
		}
	}

	return p
}

func newTariffCalc(file CellReader, layout ReceiptLayout, rowOffset int) TariffCalc {
	return TariffCalc{
		TariffName:      getCellValue(file, layout, 0, rowOffset),
		CurrentInd:      getCellValue(file, layout, 1, rowOffset),
		LastInd:         getCellValue(file, layout, 2, rowOffset),
		DifValue:        getCellValue(file, layout, 3, rowOffset),
		Ratio:           getCellValue(file, layout, 4, rowOffset),
		Step1Calc:       getCellValue(file, layout, 5, rowOffset),
		Step2Calc:       getCellValue(file, layout, 6, rowOffset),
		Step3Calc:       getCellValue(file, layout, 7, rowOffset),
		Step1Price:      getCellValue(file, layout, 8, rowOffset),
		Step2Price:      getCellValue(file, layout, 9, rowOffset),
		Step3Price:      getCellValue(file, layout, 10, rowOffset),
		Step1Arithmetic: getCellValue(file, layout, 11, rowOffset),
		Step2Arithmetic: getCellValue(file, layout, 12, rowOffset),
		Step3Arithmetic: getCellValue(file, layout, 13, rowOffset),
	}
}

func getCellValue(file CellReader, layout ReceiptLayout, difCol, difRow int) string {
	value, _ := file.GetCellValue(layout.SheetName, cellRel(layout.BaseTariffCell, difCol, difRow))
	return value
}

func cellRel(mainCell string, difCol, difRow int) string {
	col, row, _ := excelize.CellNameToCoordinates(mainCell)
	newCell, _ := excelize.CoordinatesToCellName(col+difCol, row+difRow)
	return newCell
}

func (r Receipt) SheetName() string {
	return fmt.Sprintf("%s-уч.%s", r.FullName, r.PlaceNumber)
}

func PrintSingleReceipt(f CellWriter, r Receipt) error {
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

	if err := fillSheet(f, r.SheetName(), values); err != nil {
		return err
	}

	if err := f.RemoveRow(r.SheetName(), 8); err != nil {
		return err
	}
	return nil
}

func PrintDuoReceipt(f CellWriter, r Receipt) error {
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

	return fillSheet(f, r.SheetName(), values)
}

func fillSheet(f CellWriter, sheet string, values map[string]any) error {
	for cell, val := range values {
		if err := f.SetCellValue(sheet, cell, val); err != nil {
			return err
		}
	}
	return nil
}
