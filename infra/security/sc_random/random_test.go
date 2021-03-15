package sc_random

import "testing"

func TestGenerateRandomString(t *testing.T) {
	for l := -1; l < 32; l++ {
		s, e := GetSecureRandomString(l)
		if l < 1 && e == nil {
			t.Errorf("Should fail with size (%d)", l)
		}
		if l >= 1 && (e != nil || len(s) != l) {
			t.Errorf("Error or invalid length (%d): generated (%s)", l, s)
		}
	}
}
