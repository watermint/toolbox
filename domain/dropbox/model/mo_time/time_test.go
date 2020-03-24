package mo_time

import "testing"

func TestZero(t *testing.T) {
	if !Zero().IsZero() {
		t.Error("non zero instance returned")
	}
}
