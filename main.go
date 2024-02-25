package main

import (
	"fmt"
	"go-app/application/usecase/finance"
)

func main() {
	usecase := finance.InitializeUsecase()
	result, _ := usecase.SimulateAnnualInvestment(10000, 10, 1)
	fmt.Println(result)
}
