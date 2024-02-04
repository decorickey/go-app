package investment

import "errors"
import "golang.org/x/text/message"
import "golang.org/x/text/language"

type amount struct {
	value int
}

func NewAmount(i int) (amount, error) {
	if i < 0 {
		var a amount
		return a, errors.New("invalid param")
	}
	return amount{i}, nil
}

func (a amount) Equals(v amount) bool {
	return a.value == v.value
}

func (a amount) Mult(i int) amount {
	res, _ := NewAmount(a.value * i)
	return res
}

func (a amount) String() string {
	p := message.NewPrinter(language.Japanese)
	return p.Sprint(a.value)
}

type AnnualInvestmentAmount = amount

type period struct {
	value int
}

func NewPeriod(i int) (period, error) {
	if i < 1 {
		var p period
		return p, errors.New("")
	}
	return period{i}, nil
}

type roi struct {
	value float64
}

func NewRoi(f float64) (roi, error) {
	if f < 0 {
		var r roi
		return r, errors.New("invalid param")
	}
	return roi{f}, nil
}

func Calc(a AnnualInvestmentAmount, p period, r roi) (amount, error) {
	var res float64 = 0
	for i := 1; i <= p.value; i++ {
		res += float64(a.value)
		res *= (1 + r.value/100)
	}
	return NewAmount(int(res))
}
