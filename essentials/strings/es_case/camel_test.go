package es_case

import "testing"

func TestTokenToCamelCase(t *testing.T) {
	verify(t, TokenToCamelCase, "abc", "Abc")
	verify(t, TokenToCamelCase, "ABC", "Abc")
	verify(t, TokenToCamelCase, "Abc", "Abc")
	verify(t, TokenToCamelCase, "ＡＢＣ", "Ａｂｃ")
	verify(t, TokenToCamelCase, "", "")
	verify(t, TokenToCamelCase, "PoweredByGoLang", "Poweredbygolang")
}

func TestToLowerCamelCase(t *testing.T) {
	verify(t, ToLowerCamelCase, "powered by go lang", "poweredByGoLang")
	verify(t, ToLowerCamelCase, "", "")
	verify(t, ToLowerCamelCase, "Go", "go")
	verify(t, ToLowerCamelCase, "func TestToLowerCamelCase(t *testing.T) {}", "funcTestToLowerCamelCaseTTestingT")
}

func TestToUpperCamelCase(t *testing.T) {
	verify(t, ToUpperCamelCase, "powered by go lang", "PoweredByGoLang")
	verify(t, ToUpperCamelCase, "", "")
	verify(t, ToUpperCamelCase, "Go", "Go")
	verify(t, ToUpperCamelCase, "func TestToUpperCamelCase(t *testing.T) {}", "FuncTestToUpperCamelCaseTTestingT")
}
