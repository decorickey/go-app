package service

import (
	"fmt"
	model "go-app/domain/model/investment"

	"github.com/shopspring/decimal"
)

type PeriodicInvestmentService interface {
	Simulate(monthlyInvestmentAmount int, investmentPeriod int, roi float64) (int, error)
}

func NewPeriodicInvestmentService() PeriodicInvestmentService {
	return periodicInvestmentService{}
}

type periodicInvestmentService struct{}

func (s periodicInvestmentService) Simulate(monthlyAmount int, investmentPeriod int, expectedRoi float64) (int, error) {
	amount, err := model.NewAmount(monthlyAmount)
	if err != nil {
		return 0, err
	}
	annualAmount := amount.Mul(decimal.NewFromInt(12))

	period, err := model.NewPeriod(investmentPeriod)
	if err != nil {
		return 0, err
	}

	roi, err := model.NewRoi(expectedRoi)
	if err != nil {
		return 0, err
	}

	var r decimal.Decimal
	for i := 1; i <= int(period); i++ {
		r = r.Add(annualAmount)
		r = roi.Div(decimal.NewFromInt(100)).Add(decimal.NewFromInt(1)).Mul(r)
	}
	result, err := model.NewAmount(r)
	if err != nil {
		return 0, fmt.Errorf("failed to simulate: %w", err)
	}
	return int(result.IntPart()), nil
}
