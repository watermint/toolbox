package ut_math

import "testing"

func TestMaxInt(t *testing.T) {
	if MaxInt(0, 100) != 100 {
		t.Error("invalid")
	}
	if MaxInt(100, 0) != 100 {
		t.Error("invalid")
	}
}

func TestMinInt(t *testing.T) {
	if MinInt(0, 100) != 0 {
		t.Error("invalid")
	}
	if MinInt(100, 0) != 0 {
		t.Error("invalid")
	}
}
