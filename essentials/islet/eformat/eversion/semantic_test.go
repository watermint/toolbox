package eversion

import "testing"

func TestParse(t *testing.T) {
	if v, out := Parse("1.2.3"); out.IsError() {
		t.Error(out)
	} else if v.Major != 1 || v.Minor != 2 || v.Patch != 3 || v.PreRelease != "" || v.Build != "" {
		t.Error(v)
	}

	if v, out := Parse("1.2.3-alpha"); out.IsError() {
		t.Error(out)
	} else if v.Major != 1 || v.Minor != 2 || v.Patch != 3 || v.PreRelease != "alpha" || v.Build != "" {
		t.Error(v)
	}

	if v, out := Parse("1.2.3-beta+001"); out.IsError() {
		t.Error(out)
	} else if v.Major != 1 || v.Minor != 2 || v.Patch != 3 || v.PreRelease != "beta" || v.Build != "001" {
		t.Error(v)
	}

	// invalid format
	if v, out := Parse("1-2-3-alpha"); out.IsOk() {
		t.Error(v)
	}
}

func TestCompare(t *testing.T) {
	if x := Compare(MustParse("1.0.0"), MustParse("2.0.0")); x != -1 {
		t.Error(x)
	}
	if x := Compare(MustParse("1.0.0"), MustParse("1.0.0")); x != 0 {
		t.Error(x)
	}
	if x := Compare(MustParse("2.0.0"), MustParse("1.0.0")); x != 1 {
		t.Error(x)
	}

	if x := Compare(MustParse("2.0.0"), MustParse("2.1.0")); x != -1 {
		t.Error(x)
	}
	if x := Compare(MustParse("2.1.0"), MustParse("2.0.0")); x != 1 {
		t.Error(x)
	}

	if x := Compare(MustParse("2.1.0"), MustParse("2.1.1")); x != -1 {
		t.Error(x)
	}
	if x := Compare(MustParse("2.1.1"), MustParse("2.1.0")); x != 1 {
		t.Error(x)
	}

	if x := Compare(MustParse("1.0.0-alpha"), MustParse("1.0.0")); x != -1 {
		t.Error(x)
	}
	if x := Compare(MustParse("1.0.0"), MustParse("1.0.0-alpha")); x != 1 {
		t.Error(x)
	}
}

func TestMax(t *testing.T) {
	if x := Max(MustParse("1.0.0"), MustParse("1.0.0-alpha"), MustParse("1.0.1")); !x.Equals(MustParse("1.0.1")) {
		t.Error(x)
	}
	if x := Max(MustParse("1.0.0"), MustParse("1.0.0-alpha"), MustParse("2.0.1")); !x.Equals(MustParse("2.0.1")) {
		t.Error(x)
	}
}

func TestMin(t *testing.T) {
	if x := Min(MustParse("1.0.0"), MustParse("1.0.0-alpha"), MustParse("1.0.1")); !x.Equals(MustParse("1.0.0-alpha")) {
		t.Error(x)
	}
	if x := Min(MustParse("1.0.0"), MustParse("1.0.1-alpha"), MustParse("2.0.1")); !x.Equals(MustParse("1.0.0")) {
		t.Error(x)
	}
}
