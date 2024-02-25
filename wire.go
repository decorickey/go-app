//go:build wireinject
// +build wireinject

package main

import (
	"go-app/application/domain"
	"go-app/application/usecase/finance"

	"github.com/google/wire"
)

func initializeFinanceUsecase() finance.Usecase {
	wire.Build(finance.NewUsecase, domain.NewFinanceService)
	return nil
}
