package es_case

// ToLowerKebabCase change to upper kebab case like from "powered by go-lang" to "powered-by-go-lang".
func ToLowerKebabCase(s string) string {
	return ToLowerDelimited(s, "-")
}

// ToUpperKebabCase change to upper kebab case like from "powered by go-lang" to "POWERED-BY-GO-LANG".
func ToUpperKebabCase(s string) string {
	return ToUpperDelimited(s, "-")
}
