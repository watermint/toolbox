package eg_color

import "testing"

func TestRgbaImpl_Equals(t *testing.T) {
	c1 := NewRgba(1, 2, 3, 4)
	c2 := NewRgba(4, 1, 2, 3)

	if !c1.Equals(c1) {
		t.Error(c1)
	}
	if c1.Equals(c2) {
		t.Error(c1, c2)
	}
	if c2.Equals(c1) {
		t.Error(c1, c2)
	}
	if c1.Equals(nil) {
		t.Error(c1)
	}
}
