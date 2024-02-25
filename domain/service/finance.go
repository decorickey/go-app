package service

import (
	"fmt"
	model "go-app/domain/model/finance"
)

type FinanceService interface {
	Simulate(a model.Amount, p model.Period, r model.Roi) (model.Amount, error)
}

func NewFinanceService() FinanceService {
	return financeService{}
}

type financeService struct{}

func (s financeService) Simulate(a model.Amount, p model.Period, r model.Roi) (model.Amount, error) {
	var f float64
	for i := 1; i <= p.Value; i++ {
		f += float64(a.Value)
		f *= (1 + r.Value/100)
	}
	result, err := model.NewAmount(f)
	if err != nil {
		return result, fmt.Errorf("failed to calculate annual investment amount: %w", err)
	}
	return result, nil
}
