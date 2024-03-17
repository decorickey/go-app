package investment

import "testing"

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
