package ecase

import "testing"

func verifyDelimited(t *testing.T, f func(string, string) string, s, delimiter, expected string) {
	if x := f(s, delimiter); x != expected {
		t.Error(s, delimiter, x, expected)
	}
}

func verify(t *testing.T, f func(string) string, s, expected string) {
	if x := f(s); x != expected {
		t.Error(s, x, expected)
	}
}

func TestToLowerDelimited(t *testing.T) {
	verifyDelimited(t, ToLowerDelimited, "Powered by Go-lang", "=", "powered=by=go=lang")
	verifyDelimited(t, ToLowerDelimited, "Powered by Go-lang", "", "poweredbygolang")
	verifyDelimited(t, ToLowerDelimited, "func TestToLowerDelimited(t *testing.T) {}", "=", "func=test=to=lower=delimited=t=testing=t")
	verifyDelimited(t, ToLowerDelimited, "", "=", "")
}

func TestToUpperDelimited(t *testing.T) {
	verifyDelimited(t, ToUpperDelimited, "Powered by Go-lang", "=", "POWERED=BY=GO=LANG")
	verifyDelimited(t, ToUpperDelimited, "Powered by Go-lang", "", "POWEREDBYGOLANG")
	verifyDelimited(t, ToUpperDelimited, "func TestToLowerDelimited(t *testing.T) {}", "=", "FUNC=TEST=TO=LOWER=DELIMITED=T=TESTING=T")
	verifyDelimited(t, ToUpperDelimited, "", "=", "")

}
