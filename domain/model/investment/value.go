package investment

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Amount struct {
	decimal.Decimal
}

func NewAmount(v any) (Amount, error) {
	var d decimal.Decimal
	var err error
	switch v.(type) {
	case int:
		d = decimal.NewFromInt(int64(v.(int)))
	case float64:
		d = decimal.NewFromFloat(v.(float64))
	case string:
		d, err = decimal.NewFromString(v.(string))
	case decimal.Decimal:
		d = v.(decimal.Decimal)
	}

	if err != nil {
		return Amount{}, fmt.Errorf("invalid input %v", v)
	}

	if d.IsNegative() {
		return Amount{}, fmt.Errorf("amount is zero or positive: %s", d.String())
	}

	return Amount{d}, nil
}

type Roi struct {
	decimal.Decimal
}

func NewRoi(v any) (Roi, error) {
	var d decimal.Decimal
	var err error
	switch v.(type) {
	case int:
		d = decimal.NewFromInt(int64(v.(int)))
	case float64:
		d = decimal.NewFromFloat(v.(float64))
	case string:
		d, err = decimal.NewFromString(v.(string))
	}
	if err != nil {
		return Roi{}, fmt.Errorf("invalid input %v", v)
	}

	if d.IsNegative() {
		return Roi{}, fmt.Errorf("amount is zero or positive: %s", d.String())
	}

	return Roi{d}, nil
}

type Period int

func NewPeriod(i int) (Period, error) {
	if i < 1 {
		return Period(0), fmt.Errorf("period is positive: %d", i)
	}
	return Period(i), nil
}
