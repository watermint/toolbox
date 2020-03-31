package sv_desktop

import "testing"

func TestDesktopImpl_Lookup(t *testing.T) {
	d := New()
	// No validation
	d.Lookup()
}
