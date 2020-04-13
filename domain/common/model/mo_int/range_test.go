package mo_int

import "testing"

func TestRangeInt(t *testing.T) {
	ri := NewRange()
	ri.SetRange(3, 5, 4)
	if !ri.IsValid() {
		t.Error(ri.IsValid())
	}
	if ri.Value() != 4 {
		t.Error(ri.Value())
	}
	ri.SetValue(9)
	if ri.IsValid() {
		t.Error(ri.IsValid())
	}
	if min, max := ri.Range(); min != 3 || max != 5 {
		t.Error(min, max)
	}
}
