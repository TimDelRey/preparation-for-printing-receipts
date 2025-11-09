package domain

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

// func NewPersonCalc(
// 	placeNumber int,
// 	FullName string,
// 	Summ, Balance, Payment, Debt float64,
// 	TariffName string,
// 	CurrentInd, LastInd, DifValue float64,
// 	Ratio int,
// 	Step1Calc, Step2Calc, Step3Calc float64,
// 	Step1Price, Step2Price, Step3Price float64,
// 	Step1Arithmetic, Step2Arithmetic, Step3Arithmetic float64,
// ) PersonCalc {
// 	p := PersonCalc{
// 		PlaceNumber: placeNumber,
// 		FullName:    FullName,
// 		Summ:        Summ,
// 		Balance:     Balance,
// 		Payment:     Payment,
// 		Debt:        Debt,
// 	}

// 	switch TariffName {
// 	case "Тариф 1":
// 		p.Single = &SingleZone{
// 			Tariff: TariffCalc{
// 				TariffName: TariffName,
// 				CurrentInd: CurrentInd,
// 				LastInd:    LastInd,
// 				Step1Price: Step1Price,
// 				Step2Price: Step2Price,
// 				Step3Price: Step3Price,
// 			},
// 		}
// 	default:
// 		p.Duo = &DuoCalc{
// 			T1: TariffCalc{
// 				TariffName: TariffName,
// 				CurrentInd: CurrentInd,
// 				LastInd:    LastInd,
// 				Step1Price: Step1Price,
// 				Step2Price: Step2Price,
// 				Step3Price: Step3Price,
// 			},
// 			T2: TariffCalc{
// 				TariffName: TariffName,
// 				CurrentInd: CurrentInd,
// 				LastInd:    LastInd,
// 				Step1Price: Step1Price,
// 				Step2Price: Step2Price,
// 				Step3Price: Step3Price,
// 			},
// 		}
// 	}

// 	return p
// }
