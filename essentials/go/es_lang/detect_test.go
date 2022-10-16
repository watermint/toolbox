package es_lang

import "testing"

func TestDetect(t *testing.T) {
	if c := Detect([]Lang{English}).Code(); c != "en" {
		t.Error(c)
	}
}
