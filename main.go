package main

import (
	"fmt"
)

func main() {
	usecase := initializeFinanceUsecase()
	result, _ := usecase.SimulateAnnualInvestment(10000, 10, 1)
	fmt.Println(result)
}
