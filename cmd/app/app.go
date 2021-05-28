package main

import (
	"fmt"
	"github.com/danshu93/go-skeleton/taxation"
)

func main() {
	grossIncome := taxation.CalculateGrossIncomeForDesiredNetIncome(20000)
	fmt.Println(grossIncome)
}
