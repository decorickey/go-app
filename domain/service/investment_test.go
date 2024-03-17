package service

import (
	"fmt"
	"testing"
)

func TestSimulate(t *testing.T) {
	data := []struct {
		amount   int
		period   int
		roi      float64
		expected int
	}{
		{100000, 1, 10, 1320000},
		{100000, 2, 10, 2772000},
		{100000, 10, 0, 12000000},
	}
	for i, d := range data {
		t.Run(fmt.Sprintf("No. %d", i+1), func(t *testing.T) {
			service := NewPeriodicInvestmentService()
			actual, _ := service.Simulate(d.amount, d.period, d.roi)
			if actual != d.expected {
				t.Errorf("expected: %v, actual: %v", d.expected, actual)
			}
		})
	}
}
