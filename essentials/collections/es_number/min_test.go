package es_number

import "testing"

func TestMin(t *testing.T) {
	// int sum
	if s := Min(1, 7, 3, 6, 5, 4).Int(); s != 1 {
		t.Error(s)
	}

	// float
	if s := Min(1.0, 7.0, 3.0, 6.0, 5.0, 4.0).Float64(); s != 1.0 {
		t.Error(s)
	}
}
