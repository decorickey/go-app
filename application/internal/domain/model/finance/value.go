package finance

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Amount struct {
	Value int
}

func NewAmount[T int | float64](n T) (Amount, error) {
	i := int(n)
	if i < 0 {
		return Amount{}, fmt.Errorf("failed to create amount")
	}
	return Amount{i}, nil
}

func (a Amount) Equals(v Amount) bool {
	return a.Value == v.Value
}

func (a Amount) Mult(i int) Amount {
	res, _ := NewAmount(a.Value * i)
	return res
}

func (a Amount) String() string {
	p := message.NewPrinter(language.Japanese)
	return p.Sprint(a.Value)
}

type Period struct {
	Value int
}

func NewPeriod(i int) (Period, error) {
	if i < 1 {
		var p Period
		return p, fmt.Errorf("failed to create period")
	}
	return Period{i}, nil
}

type Roi struct {
	Value float64
}

func NewRoi(f float64) (Roi, error) {
	if f < 0 {
		var r Roi
		return r, fmt.Errorf("failed to create period")
	}
	return Roi{f}, nil
}
