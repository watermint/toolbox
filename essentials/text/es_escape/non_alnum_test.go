package es_escape

import "testing"

func TestReplaceNonAlNum(t *testing.T) {
	if x := ReplaceNonAlNum("abc-DEF-012_[4]", "#"); x != "abc#DEF#012##4#" {
		t.Error(x)
	}
	if x := ReplaceNonAlNum("abc", "#"); x != "abc" {
		t.Error(x)
	}
	if x := ReplaceNonAlNum("ABC", "#"); x != "ABC" {
		t.Error(x)
	}
}
