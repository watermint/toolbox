package number

import "testing"

func TestSum(t *testing.T) {
	// int sum
	if s := Sum(1, 2, 3, 4, 5, 6, 7).Int(); s != 28 {
		t.Error(s)
	}

	// float
	if s := Sum(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0).Float64(); s != 28 {
		t.Error(s)
	}
}
