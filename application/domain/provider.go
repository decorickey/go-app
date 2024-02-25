package domain

import "go-app/application/internal/domain/service"

func NewFinanceService() service.FinanceService {
	return service.NewFinanceService()
}
