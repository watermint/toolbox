package mo_filter

import "testing"

func TestFilterImpl_Accept(t *testing.T) {
	fl := New("hello")

	// always accept
	if x := fl.Accept(123); !x {
		t.Error(x)
	}
	if fl.Name() != "hello" {
		t.Error(fl.Name())
	}
	fl.SetOptions(NewNameFilter(), NewNameSuffixFilter())
}
