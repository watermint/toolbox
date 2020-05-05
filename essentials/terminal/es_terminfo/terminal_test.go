package es_terminfo

import "testing"

func TestIsTerminal(t *testing.T) {
	// ensure the func should not panic
	_ = IsInTerminal()
	_ = IsOutTerminal()
	_ = IsOutColorTerminal()
}
