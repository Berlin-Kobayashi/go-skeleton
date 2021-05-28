package taxation

import "testing"

type taxCase struct {
	description                   string
	desiredNetIncome, grossIncome int
}

var taxCases = []taxCase{
	{"Taxless", 1500, 1500},
	{"First progression", 10000, 10000},
	{"Second progression", 20000, 20000},
	{"Third progression", 100000, 100000},
	{"Fourth progression", 500000, 500000},
}

func TestCalculateGrossIncomeForDesiredNetIncome(t *testing.T) {
	for _, taxCase := range taxCases {
		actual := CalculateGrossIncomeForDesiredNetIncome(taxCase.desiredNetIncome)
		if actual != taxCase.grossIncome {
			t.Errorf(
				"necessary gross income did not match for case %q expected value of %d, got %d",
				taxCase.description,
				taxCase.grossIncome,
				actual,
			)
		}
	}
}
