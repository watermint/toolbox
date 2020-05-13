package mo_filter

import (
	"github.com/watermint/toolbox/essentials/collections/es_array"
	"testing"
)

func TestFilterImpl_Accept(t *testing.T) {
	fl := New("Hello")

	// always accept
	if x := fl.Accept(123); !x {
		t.Error(x)
	}
	if fl.Name() != "Hello" {
		t.Error(fl.Name())
	}
	fl.SetOptions(NewNameFilter(), NewNameSuffixFilter())
	fields := es_array.NewByString(fl.Fields()...)
	expected := es_array.NewByString("HelloName", "HelloNameSuffix")
	if x := fields.Intersection(expected); x.Size() != 2 {
		t.Error(x)
	}
}
