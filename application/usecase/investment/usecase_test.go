package investment

import (
	"fmt"
	"go-app/domain/service"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := service.NewMockPeriodicInvestmentService(ctrl)
	usecase := NewUsecase(service)

	data := []struct {
		amount      int
		period      int
		roi         float64
		expectedVal int
		expectedErr error
	}{
		{1, 1, 1, 1, nil},
	}

	for i, d := range data {
		t.Run(fmt.Sprintf("No. %d", i+1), func(t *testing.T) {
			service.EXPECT().Simulate(gomock.Any(), gomock.Any(), gomock.Any()).Return(d.expectedVal, d.expectedErr)
			actualVal, actualErr := usecase.SimulatePeriodicInvestment(d.amount, d.period, d.roi)
			if actualVal != d.expectedVal {
				t.Errorf("expected: %v, actual: %v", d.expectedVal, actualVal)
			}
			if actualErr != d.expectedErr {
				t.Errorf("expected: %v, actual: %v", d.expectedErr, actualErr)
			}
		})
	}

}
