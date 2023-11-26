package es_case

// ToLowerSnakeCase change to snake case like from "Powered by go-lang" to "powered_by_go_lang".
func ToLowerSnakeCase(s string) string {
	return ToLowerDelimited(s, "_")
}

// ToUpperSnakeCase change to snake case like from "Powered by go-lang" to "POWERED_BY_GO_LANG".
func ToUpperSnakeCase(s string) string {
	return ToUpperDelimited(s, "_")
}
