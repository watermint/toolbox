package mo_filter

import (
	"reflect"
	"slices"
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
	if x := fl.IsEnabled(); x {
		t.Error(x)
	}
	fl.SetOptions(NewNameFilter(), NewNameSuffixFilter())
	fields := fl.Fields()
	expected := []string{"HelloName", "HelloNameSuffix"}

	// Sort both slices for comparison
	slices.Sort(fields)
	slices.Sort(expected)

	if !reflect.DeepEqual(fields, expected) {
		t.Errorf("Expected fields %v but got %v", expected, fields)
	}
}
