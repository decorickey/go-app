package finance

import (
	"fmt"
	"go-app/domain/model/finance"
	"go-app/domain/service"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := service.NewMockFinanceService(ctrl)
	usecase := NewUsecase(service)

	data := []struct {
		amount      int
		period      int
		roi         float64
		expectedVal finance.Amount
		expectedErr error
	}{
		{1, 1, 1, finance.Amount{Value: 1}, nil},
	}

	for i, d := range data {
		t.Run(fmt.Sprintf("No. %d", i+1), func(t *testing.T) {
			service.EXPECT().SimulateAnnualInvestment(gomock.Any(), gomock.Any(), gomock.Any()).Return(d.expectedVal, d.expectedErr)
			actualVal, actualErr := usecase.SimulateAnnualInvestment(d.amount, d.period, d.roi)
			if actualVal != d.expectedVal.Value {
				t.Errorf("expected: %v, actual: %v", d.expectedVal, actualVal)
			}
			if actualErr != d.expectedErr {
				t.Errorf("expected: %v, actual: %v", d.expectedErr, actualErr)
			}
		})
	}

}
