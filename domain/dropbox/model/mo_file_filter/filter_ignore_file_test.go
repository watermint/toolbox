package mo_file_filter

import "testing"

func TestFilterIgnoreFileOpt_Accept(t *testing.T) {
	op := NewIgnoreFileFilter()
	if x := op.Enabled(); !x {
		t.Error(x)
	}
	b := op.Bind().(*bool)
	*b = true
	if x := op.Enabled(); x {
		t.Error(x)
	}
	*b = false

	if x := op.NameSuffix(); x != "DisableIgnore" {
		t.Error(x)
	}
	if x := op.Desc(); x.Key() != MFileFilterOpt.Desc.Key() {
		t.Error(x)
	}

	if x := op.Accept(".dropbox"); x {
		t.Error(x)
	}
	if x := op.Accept("watermint"); !x {
		t.Error(x)
	}
}
