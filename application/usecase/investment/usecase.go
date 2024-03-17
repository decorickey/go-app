package investment

import (
	"fmt"

	"go-app/domain/service"
)

type Usecase interface {
	SimulatePeriodicInvestment(a, p int, r float64) (int, error)
}

func NewUsecase(service service.PeriodicInvestmentService) Usecase {
	return usecase{service}
}

type usecase struct {
	service service.PeriodicInvestmentService
}

func (u usecase) SimulatePeriodicInvestment(a, p int, r float64) (int, error) {
	result, err := u.service.Simulate(a, p, r)
	if err != nil {
		return 0, fmt.Errorf("unexpected error: %w", err)
	}

	return result, nil
}
