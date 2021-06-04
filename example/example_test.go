package example

import "testing"

type exampleCase struct {
	description           string
	input, expectedOutput int
}

var exampleCases = []exampleCase{
	{"Positive", 3, 6},
	{"Negative", -4, -8},
	{"Zero", 0, 0},
}

func TestCalculateGrossIncomeForDesiredNetIncome(t *testing.T) {
	for _, exampleCase := range exampleCases {
		actual := Example(exampleCase.input)
		if actual != exampleCase.expectedOutput {
			t.Errorf(
				"Example did not match for case %q expected value of %d, got %d",
				exampleCase.description,
				exampleCase.expectedOutput,
				actual,
			)
		}
	}
}
