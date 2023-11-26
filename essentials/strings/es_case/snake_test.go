package es_case

import "testing"

func TestToLowerSnakeCase(t *testing.T) {
	verify(t, ToLowerSnakeCase, "powered by go lang", "powered_by_go_lang")
	verify(t, ToLowerSnakeCase, "", "")
	verify(t, ToLowerSnakeCase, "Go", "go")
	verify(t, ToLowerSnakeCase, "func TestToLowerSnakeCase(t *testing.T) {}", "func_test_to_lower_snake_case_t_testing_t")
}

func TestToUpperSnakeCase(t *testing.T) {
	verify(t, ToUpperSnakeCase, "powered by go lang", "POWERED_BY_GO_LANG")
	verify(t, ToUpperSnakeCase, "", "")
	verify(t, ToUpperSnakeCase, "Go", "GO")
	verify(t, ToUpperSnakeCase, "func TestToUpperSnakeCase(t *testing.T) {}", "FUNC_TEST_TO_UPPER_SNAKE_CASE_T_TESTING_T")
}
