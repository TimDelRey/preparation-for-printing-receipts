package samples

import "strings"

var (
	CalcSingleHeaders = map[string]string{
		// шапка
		"B1:M1": `Организация: Товарищество собственников недвижимости "жилищно-строительное товарищество индивидуальных застройщиков "Столетовский"`,
		"B2:M2": `ИНН: 9201505978; КПП: 920101001; ФИЛИАЛ "ЦЕНТРАЛЬНЫЙ" БАНКА ВТБ (ПАО) в г. Москве;`,
		"B3:M3": `БИК 044525411; Р/с 40703810824317000026; К/с 30101810145250000411`,
		"B4:M4": `тел. +7(978) 812-13-99, e-mail: tsn-gstiz-stoletovskiy@yandex.ru`,
		"R1:S1": `дата`,
		// тело
		"B5:D5": `г. Севастополь,
ТСН "ЖСТИЗ "СТОЛЕТОВСКИЙ`,
		"E5:G5": `Показания по прибору учета электроэнергии`,
		"H5:J5": `Расчет обьема,
кВт.*ч`,
		"K5:M5": `Ставка тарифа,
р./кВт.*ч`,
		"N5:P5": `Начислено, руб.`,
		"B6":    `Уч.№`,
		"C6":    `Потребитель`,
		"D6":    `Тарифная зона`,
		"E6": `Текущий
месяц, кВт.*ч`,
		"F6": `Предыдущий
месяц, кВт.*ч`,
		"G6": `Расход,
кВт*ч`,
		"H6": `Блок 1`,
		"I6": `Блок 2`,
		"J6": `Блок 3`,
		"K6": `Блок 1`,
		"L6": `Блок 2`,
		"M6": `Блок 3`,
		"N6": `Блок 1`,
		"O6": `Блок 2`,
		"P6": `Блок 3`,
		"Q5:Q6": `Сумма,
руб.`,
		"R5:R6": `Долг / Переплата
"-", руб.`,
		"S5:S6": `Учтена оплата, руб.`,
		"T5:T6": `К оплате, руб.`,
		// подвал
		"B8:T8": `По решению общего собрания потери в электросети ТСН по итогам года рассчитываются равным образом между потребителями.
Согласно приложению к приказу №185-УТ от 09.12.2024 г. Управления по тарифам города Севастополя на электроэнергию для населения и приравненных к нему категорий потребителей СНТ в г. Севастополе:
- в одноставочном тарифе в I-полугодие/II-полугодие, объем, потребляемый до 150 кВт*ч составляет 4,49/5,05 руб./кВт*ч; объем, потребляемый от 151 кВт*ч до 600 кВт*ч – 4,82/5,47 руб./кВт*ч; объем более 600 кВт*ч – 8,28/9,4 руб./кВт*ч;
- в дифференцированном по двум зонам суток, дневной тариф на электрическую энергию, в I-полугодие/II-полугодие составляет: объем,
потребляемый до 150 кВт*ч – 4,49/5,05 руб./кВт*ч; объем от 151 кВт*ч до 600 кВт*ч – 4,82/5,47 руб./кВт*ч; объем более 600 кВт*ч – 8,28/9,4 руб./кВт*ч.
- в дифференцированном по двум зонам суток, ночной тариф на электрическую энергию, в I-полугодие/II-полугодие составляет: объем,
потребляемый до 150 кВт*ч – 3,13/3,52 руб./кВт*ч; объем от 151 кВт*ч до 600 кВт*ч– 3,74/4,21 руб./кВт*ч; объем более 600 кВт*ч – 6,21/6,99 руб./кВт*ч.`,
	}

	CalcDuoHeaders = map[string]string{
		// хедер
		"B1:M1": `Организация: Товарищество собственников недвижимости "жилищно-строительное товарищество индивидуальных застройщиков "Столетовский"`,
		"B2:M2": `ИНН: 9201505978; КПП: 920101001; ФИЛИАЛ "ЦЕНТРАЛЬНЫЙ" БАНКА ВТБ (ПАО) в г. Москве;`,
		"B3:M3": `БИК 044525411; Р/с 40703810824317000026; К/с 30101810145250000411`,
		"B4:M4": `тел. +7(978) 812-13-99, e-mail: tsn-gstiz-stoletovskiy@yandex.ru`,
		"S1:T1": `дата`,
		// тело
		"B5:D5": `г. Севастополь,
ТСН "ЖСТИЗ "СТОЛЕТОВСКИЙ`,
		"E5:H5": `Показания по прибору учета электроэнергии`,
		"I5:K5": `Расчет обьема,
кВт.*ч`,
		"L5:N5": `Ставка тарифа,
р./кВт.*ч`,
		"O5:Q5": `Начислено, руб.`,
		"B6":    `Уч.№`,
		"C6":    `Потребитель`,
		"D6":    `Тарифная зона`,
		"E6": `Текущий
месяц, кВт.*ч`,
		"F6": `Предыдущий
месяц, кВт.*ч`,
		"G6": `Расход,
кВт*ч`,
		"H6": `Соотно-
шение, %`,
		"I6": `Блок 1`,
		"J6": `Блок 2`,
		"K6": `Блок 3`,
		"L6": `Блок 1`,
		"M6": `Блок 2`,
		"N6": `Блок 3`,
		"O6": `Блок 1`,
		"P6": `Блок 2`,
		"Q6": `Блок 3`,
		"R5:R6": `Сумма,
руб.`,
		"S5:S6": `Долг / Переплата
"-", руб.`,
		"T5:T6": `Учтена оплата, руб.`,
		"U5:U6": `К оплате, руб.`,
		// подвал
		"B9:T9": `По решению общего собрания потери в электросети ТСН по итогам года рассчитываются равным образом между потребителями.
Согласно приложению к приказу №185-УТ от 09.12.2024 г. Управления по тарифам города Севастополя на электроэнергию для населения и приравненных к нему категорий потребителей СНТ в г. Севастополе:
- в одноставочном тарифе в I-полугодие/II-полугодие, объем, потребляемый до 150 кВт*ч составляет 4,49/5,05 руб./кВт*ч; объем, потребляемый от 151 кВт*ч до 600 кВт*ч – 4,82/5,47 руб./кВт*ч; объем более 600 кВт*ч – 8,28/9,4 руб./кВт*ч;
- в дифференцированном по двум зонам суток, дневной тариф на электрическую энергию, в I-полугодие/II-полугодие составляет: объем,
потребляемый до 150 кВт*ч – 4,49/5,05 руб./кВт*ч; объем от 151 кВт*ч до 600 кВт*ч – 4,82/5,47 руб./кВт*ч; объем более 600 кВт*ч – 8,28/9,4 руб./кВт*ч.
- в дифференцированном по двум зонам суток, ночной тариф на электрическую энергию, в I-полугодие/II-полугодие составляет: объем,
потребляемый до 150 кВт*ч – 3,13/3,52 руб./кВт*ч; объем от 151 кВт*ч до 600 кВт*ч– 3,74/4,21 руб./кВт*ч; объем более 600 кВт*ч – 6,21/6,99 руб./кВт*ч.`,
	}

	calcSingleWidth = map[string]float64{
		"A": 2.01,
		"B": 4.51,
		"C": 15.01,
		"D": 7.01,
		"E": 8.34,
		"F": 8.67,
		"G": 6.84,
		"H": 6.51,
		"I": 6.51,
		"J": 6.51,
		"K": 5.01,
		"L": 5.01,
		"M": 5.01,
		"N": 6.51,
		"O": 6.51,
		"P": 6.51,
		"Q": 7.17,
		"R": 8.67,
		"S": 8.67,
		"T": 7.17,
		"U": 2.01,
	}
	calcHeight = map[int]float64{
		1: 9.75,
		2: 9.75,
		3: 9.75,
		4: 9.75,
		5: 21.75,
		6: 30.75,
		7: 9.75,
		8: 9.75,
		9: 87.75,
	}
	CalcDuoWidth = map[string]float64{
		"A": 2.01,
		"B": 4.51,
		"C": 15.01,
		"D": 7.01,
		"E": 8.34,
		"F": 8.67,
		"G": 6.84,
		"H": 3.34,
		"I": 6.51,
		"J": 6.51,
		"K": 6.51,
		"L": 5.01,
		"M": 5.01,
		"N": 5.01,
		"O": 6.51,
		"P": 6.51,
		"Q": 6.51,
		"R": 7.17,
		"S": 8.67,
		"T": 8.67,
		"U": 7.17,
		"V": 2.01,
	}
)

func newSingleSample(f *excelize.File, name string) err {
	page, err := f.NewSheet(name)
	if err != nil {
		return err
	}

	file.SetActiveSheet(page)
	// установка ширины ячеек
	if err := setCellsWidth(f, calcSingleWidth); err != nil {
		return err
	}
	// установка высоты ячеек
	if err := setCellsHeight(f, calcHeight); err != nil {
		return err
	}
	// установка шрифта, размераотрисовка границ
	if err := setSingleStyle(f); err != nil {
		return err
	}
	// обьединение ячеек
	if err := mergeCells(f, calcSingleHeaders); err != nil {
		return err
	}
	// установка хедеров и подвала
	if err := setHeaders(f, calcSingleHeaders); err != nil {
		return err
	}
	return nil
}

func newDuoSample() {
	page, err := f.NewSheet(name)
	if err != nil {
		return err
	}

	file.SetActiveSheet(page)
	// установка ширины ячеек
	if err := setCellsWidth(f, calcDuoWidth); err != nil {
		return err
	}
	// установка высоты ячеек
	if err := setCellsHeight(f, calcHeight); err != nil {
		return err
	}
	// установка шрифта, размераотрисовка границ
	if err := setDuoStyle(f); err != nil {
		return err
	}
	// обьединение ячеек
	if err := mergeCells(f, calcSingleHeaders); err != nil {
		return err
	}
	// установка хедеров и подвала
	if err := setHeaders(f, calcSingleHeaders); err != nil {
		return err
	}
	return nil

}

func setCellsHeight(f *excelize.File, h map[int]float64) error {
	for i, v := range h {
		if err := f.SetRowHeight(sheet, i, v); err != nil {
			return err
		}
	}
	return nil
}

func setCellsHeight(f *excelize.File, h map[int]float64) error {
	for i, v := range h {
		if err := f.SetRowHeight(sheet, i, v); err != nil {
			return err
		}
	}
	return nil
}
func setSingleStyle(f *excelize.File) error {
	if err := setMainStyle(f, "B5", "T7"); err != nil {
		return err
	}
	if err := setNameStyle(f, "B7", "D7"); err != nil {
		return err
	}
	if err := setValueStyle(f, "E7", "T7"); err != nil {
		return err
	}
	return nil
}
func setMainStyle(f *excelize.File, from, to string) error {
	style, err := f.NewStyle(baseStyle(7, "center", true))
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, from, to, style)
}
func setNameStyle(f *excelize.File, from, to string) error {
	style, err := f.NewStyle(baseStyle(7, "left", false))
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, from, to, style)
}
func setValueStyle(f *excelize.File, from, to string) error {
	style, err := f.NewStyle(baseStyle(7, "right", true))
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, from, to, style)
}
func baseStyle(fontSize float64, horizontal string, wrap bool) *excelize.Style {
	return &excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
			{Type: "horizontal", Color: "000000", Style: 1},
			{Type: "vertical", Color: "000000", Style: 1},
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
func mergeCells(f *excelize.File, h map[string]string) error {
	for i := range h {
		if err := mergeCell(f, i); err != nil {
			return err
		}
	}
	return nil
}
func mergeCell(f *excelize.File, in string) error {
	parts := strings.Split(in, ":")
	if len(parts) == 2 {
		if err := f.MergeCell(sheet, parts[0], parts[1]); err != nil {
			return err
		}
	}
	return nil
}
func setHeaders(f *excelize.File, h map[string]string) error {
	for i, v := range h {
		parts := strings.Split(i, ":")
		if err := f.SetCellValue(sheet, parts[0], v); err != nil {
			return err
		}
	}
	return nil
}
func setDuoStyle(f *excelize.File) error {
	if err := setMainStyle(f, "B5", "U8"); err != nil {
		return err
	}
	if err := setNameStyle(f, "B7", "D8"); err != nil {
		return err
	}
	if err := setRatioStyle(f, "H6", "H6"); err != nil {
		return err
	}
	if err := setValueStyle(f, "E7", "U8"); err != nil {
		return err
	}
	if err := mergeCells(f, map[string]string{
		"B7:B8": "участок",
		"C7:C8": "имя",
		"R7:R8": "сумма",
		"S7:S8": "баланс",
		"T7:T8": "оплата",
		"U7:U8": "к оплате",
	}); err != nil {
		return err
	}
	return nil
}
func setRatioStyle(f *excelize.File, from, to string) error {
	s := baseStyle(6, "center", true)
	s.Alignment.TextRotation = 90
	style, err := f.NewStyle(s)
	if err != nil {
		return err
	}
	return f.SetCellStyle(sheet, from, to, style)
}
