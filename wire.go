//go:build wireinject
// +build wireinject

package main

import (
	"go-app/application/usecase/finance"
	"go-app/domain/service"

	"github.com/google/wire"
)

func initializeFinanceUsecase() finance.Usecase {
	wire.Build(
		finance.NewUsecase,
		service.NewFinanceService,
	)
	return nil
}
