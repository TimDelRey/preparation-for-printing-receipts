package domain

import "fmt"

type PersonCalc struct {
	PlaceNumber int
	FullName    string

	Single *SingleZone
	Duo    *DuoCalc

	Summ    float64
	Balance float64
	Payment float64
	Debt    float64
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
	CurrentInd      float64
	LastInd         float64
	DifValue        float64
	Ratio           int
	Step1Calc       float64
	Step2Calc       float64
	Step3Calc       float64
	Step1Price      float64
	Step2Price      float64
	Step3Price      float64
	Step1Arithmetic float64
	Step2Arithmetic float64
	Step3Arithmetic float64
}

func NewPersonCalc(u User, t Tariff, current, prev Indication) PersonCalc {
	p := PersonCalc{
		PlaceNumber: u.PlaceNumber,
		FullName:    fmt.Sprintf("%s %s %s", u.FirstName, u.Name, u.LastName),
	}

	switch u.Tariff {
	case "mono":
		p.Single = &SingleZone{
			Tariff: TariffCalc{
				TariffName: "Одноставочный",
				CurrentInd: current.AllDayData,
				LastInd:    prev.AllDayData,
				Step1Price: t.FirstValue,
				Step2Price: t.SecondValue,
				Step3Price: t.ThirdValue,
			},
		}
	case "duo":
		p.Duo = &DuoCalc{
			T1: TariffCalc{
				TariffName: "Дневной Т1",
				CurrentInd: current.AllDayData,
				LastInd:    prev.AllDayData,
				Step1Price: t.FirstValue,
				Step2Price: t.SecondValue,
				Step3Price: t.ThirdValue,
			},
			T2: TariffCalc{
				TariffName: "Ночной Т2",
				CurrentInd: current.AllDayData,
				LastInd:    prev.AllDayData,
				Step1Price: t.FirstValue,
				Step2Price: t.SecondValue,
				Step3Price: t.ThirdValue,
			},
		}
	}

	return p
}
