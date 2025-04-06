package mo_filter

import (
	"reflect"
	"sort"
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
	sort.Strings(fields)
	sort.Strings(expected)

	if !reflect.DeepEqual(fields, expected) {
		t.Errorf("Expected fields %v but got %v", expected, fields)
	}
}
