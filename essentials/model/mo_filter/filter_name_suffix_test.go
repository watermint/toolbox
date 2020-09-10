package mo_filter

import "testing"

func TestNameFilterSuffixOpt(t *testing.T) {
	fo := NewNameSuffixFilter()
	// should be false before flag set
	if x := fo.Enabled(); x {
		t.Error(x)
	}
	if x := fo.Desc(); x.Key() != MFilter.DescFilterNameSuffix.Key() {
		t.Error(x)
	}
	if x := fo.NameSuffix(); x != "NameSuffix" {
		t.Error(x)
	}
	b := fo.Bind()
	b0 := b.(*string)
	*b0 = "hello"

	if x := fo.Accept("helloWorld"); x {
		t.Error(b, x)
	}
	if x := fo.Accept("World_hello"); !x {
		t.Error(b, x)
	}
	if x := fo.Enabled(); !x {
		t.Error(x)
	}
}
