package mo_filter

import "testing"

func TestNameFilterPrefixOpt(t *testing.T) {
	fo := NewNamePrefixFilter()
	// should be false before flag set
	if x := fo.Enabled(); x {
		t.Error(x)
	}
	if x := fo.Desc(); x.Key() != MFilter.DescFilterNamePrefix.Key() {
		t.Error(x)
	}
	if x := fo.NameSuffix(); x != "NamePrefix" {
		t.Error(x)
	}
	b := fo.Bind()
	b0 := b.(*string)
	*b0 = "hello"

	if x := fo.Accept("helloWorld"); !x {
		t.Error(b, x)
	}
	if x := fo.Accept("worldHello"); x {
		t.Error(b, x)
	}
	if x := fo.Enabled(); !x {
		t.Error(x)
	}
}
