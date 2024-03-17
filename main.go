package main

import (
	"fmt"
)

func main() {
	// s := service.NewPeriodicInvestmentService()
	// u := investment.NewUsecase(s)
	u := initializeInvestmentUsecase()
	result, _ := u.SimulatePeriodicInvestment(100000, 30, 5)
	fmt.Println(result)
}
