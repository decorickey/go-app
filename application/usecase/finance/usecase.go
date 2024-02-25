package finance

import (
	"fmt"

	model "go-app/application/internal/domain/model/finance"
	"go-app/application/internal/domain/service"
)

type Usecase interface {
	SimulateAnnualInvestment(a, p int, r float64) (int, error)
}

func NewUsecase(service service.FinanceService) Usecase {
	return usecase{service}
}

type usecase struct {
	service service.FinanceService
}

func (u usecase) SimulateAnnualInvestment(a, p int, r float64) (int, error) {
	amount, err := model.NewAmount(a)
	if err != nil {
		return 0, fmt.Errorf("invalid annual investment amount: %w:", err)
	}
	period, err := model.NewPeriod(p)
	if err != nil {
		return 0, fmt.Errorf("invalid period: %w:", err)
	}
	roi, err := model.NewRoi(r)
	if err != nil {
		return 0, fmt.Errorf("invalid roi: %w:", err)
	}

	result, err := u.service.Simulate(amount, period, roi)
	if err != nil {
		return 0, fmt.Errorf("unexpected error: %w", err)
	}

	return result.Value, nil
}
