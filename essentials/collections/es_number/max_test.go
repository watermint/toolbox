package es_number

import "testing"

func TestMax(t *testing.T) {
	// int sum
	if s := Max(1, 7, 3, 6, 5, 4).Int(); s != 7 {
		t.Error(s)
	}

	// float
	if s := Max(1.0, 7.0, 3.0, 6.0, 5.0, 4.0).Float64(); s != 7.0 {
		t.Error(s)
	}
}
