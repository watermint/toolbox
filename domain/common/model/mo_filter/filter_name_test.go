package mo_filter

import "testing"

func TestNameFilterOpt(t *testing.T) {
	fo := NewNameFilter()
	// should be false before flag set
	if x := fo.Enabled(); x {
		t.Error(x)
	}
	if x := fo.Desc(); x.Key() != MFilter.DescFilterName.Key() {
		t.Error(x)
	}
	if x := fo.NameSuffix(); x != "Name" {
		t.Error(x)
	}
	b := fo.Bind()
	b0 := b.(*string)
	*b0 = "hello"

	if x := fo.Accept("hello"); !x {
		t.Error(b, x)
	}
	if x := fo.Accept("world"); x {
		t.Error(b, x)
	}
	if x := fo.Enabled(); !x {
		t.Error(x)
	}
}
