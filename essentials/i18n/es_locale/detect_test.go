package es_locale

import "testing"

func TestCurrentLocale(t *testing.T) {
	if l := CurrentLocale(); l.Language() == "" {
		t.Error(l.String())
	}
}
