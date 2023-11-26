package ecase

import "testing"

func TestToLowerKebabCase(t *testing.T) {
	verify(t, ToLowerKebabCase, "powered by go lang", "powered-by-go-lang")
	verify(t, ToLowerKebabCase, "", "")
	verify(t, ToLowerKebabCase, "Go", "go")
	verify(t, ToLowerKebabCase, "func TestToLowerKebabCase(t *testing.T) {}", "func-test-to-lower-kebab-case-t-testing-t")
}

func TestToUpperKebabCase(t *testing.T) {
	verify(t, ToUpperKebabCase, "powered by go lang", "POWERED-BY-GO-LANG")
	verify(t, ToUpperKebabCase, "", "")
	verify(t, ToUpperKebabCase, "Go", "GO")
	verify(t, ToUpperKebabCase, "func TestToUpperKebabCase(t *testing.T) {}", "FUNC-TEST-TO-UPPER-KEBAB-CASE-T-TESTING-T")
}
