//go:build wireinject
// +build wireinject

package finance

import (
	"go-app/application/internal/domain/service"

	"github.com/google/wire"
)

func InitializeUsecase() Usecase {
	wire.Build(NewUsecase, service.NewFinanceService)
	return nil
}
