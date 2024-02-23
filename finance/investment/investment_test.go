package investment

import (
	"testing"
)

func TestNewAmount(t *testing.T) {
	_, err := NewAmount(-1)
	if err == nil {
		t.Errorf("expected: error, actual: success")
	}

	_, err = NewAmount(0)
	if err != nil {
		t.Errorf("expected success, actual: error")
	}
}

func TestNewPeriod(t *testing.T) {
	_, err := NewPeriod(0)
	if err == nil {
		t.Errorf("expected: error, actual: success")
	}

	_, err = NewPeriod(1)
	if err != nil {
		t.Errorf("expected success, actual: error")
	}
}

func TestNewRoi(t *testing.T) {
	_, err := NewRoi(-1)
	if err == nil {
		t.Errorf("expected: error, actual: success")
	}

	_, err = NewRoi(0)
	if err != nil {
		t.Errorf("expected success, actual: error")
	}

	_, err = NewRoi(1.1)
	if err != nil {
		t.Errorf("expected success, actual: error")
	}
}

func TestCalc(t *testing.T) {
	a, _ := NewAmount(1000000)
	p, _ := NewPeriod(1)
	r, _ := NewRoi(10)
	expected, _ := NewAmount(1100000)
	actual, _ := Calc(a, p, r)
	if !actual.Equals(expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	a, _ = NewAmount(1000000)
	p, _ = NewPeriod(2)
	r, _ = NewRoi(10)
	actual, _ = Calc(a, p, r)
	expected, _ = NewAmount(2310000)
	if !actual.Equals(expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}

	a, _ = NewAmount(1000000)
	p, _ = NewPeriod(10)
	r, _ = NewRoi(0)
	actual, _ = Calc(a, p, r)
	if !actual.Equals(a.Mult(p.value)) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
