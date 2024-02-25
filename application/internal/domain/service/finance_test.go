package service

import (
	"fmt"
	model "go-app/application/internal/domain/model/finance"
	"testing"
)

func TestCalc(t *testing.T) {
	data := []struct {
		amount   int
		period   int
		roi      float64
		expected int
	}{
		{1000000, 1, 10, 1100000},
		{1000000, 2, 10, 2310000},
		{1000000, 10, 0, 10000000},
	}
	for i, d := range data {
		t.Run(fmt.Sprintf("No. %d", i+1), func(t *testing.T) {
			a, _ := model.NewAmount(d.amount)
			p, _ := model.NewPeriod(d.period)
			r, _ := model.NewRoi(d.roi)
			expected, _ := model.NewAmount(d.expected)
			service := financeService{}
			actual, _ := service.Simulate(a, p, r)
			if !actual.Equals(expected) {
				t.Errorf("expected: %v, actual: %v", expected, actual)
			}
		})
	}
}
