package mo_string

import "testing"

func TestOptString(t *testing.T) {
	s1 := NewOptional("")
	if s1.IsExists() {
		t.Error(s1.IsExists())
	}
	if s1.Value() != "" {
		t.Error(s1.Value())
	}
	s2 := NewOptional("s2")
	if !s2.IsExists() {
		t.Error(s2.IsExists())
	}
	if s2.Value() != "s2" {
		t.Error(s2.Value())
	}
}
