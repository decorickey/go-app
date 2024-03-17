//go:build wireinject
// +build wireinject

package main

import (
	"go-app/application/usecase/investment"
	"go-app/domain/service"

	"github.com/google/wire"
)

func initializeInvestmentUsecase() investment.Usecase {
	wire.Build(
		service.NewPeriodicInvestmentService,
		investment.NewUsecase,
	)
	return nil
}
