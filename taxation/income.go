package taxation

func CalculateGrossIncomeForDesiredNetIncome(desiredNetIncome int) int {
	if desiredNetIncome <= 9408 {
		return desiredNetIncome
	}

	// 10000 = (972,87 * ((out - 9408)/10000) + 1400) * ((out - 9408)/10000)
	// 10000 = (972,87 * ((out - 9408)/10000) + 1400) * ((out - 9408)/10000)

	return int((float64(desiredNetIncome) + 17078.74) / 0.45)
}
