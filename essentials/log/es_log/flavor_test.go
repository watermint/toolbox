package es_log

import "testing"

func TestNewFlavor(t *testing.T) {
	// ensure should not panic
	_ = newFlavor(FlavorConsole)
	_ = newFlavor(FlavorFileStandard)
	_ = newFlavor(FlavorFileCapture)
}
