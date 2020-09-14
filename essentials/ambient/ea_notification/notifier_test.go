package ea_notification

import "testing"

func TestRepoImpl_OnProgress(t *testing.T) {
	ok := false
	Global().OnProgress(func() {
		ok = true
	})
	if !ok {
		t.Error(ok)
	}

	Global().Suppress()
	Global().OnProgress(func() {
		t.Error("should be suppressed")
	})
	Global().Resume()

	ok = false
	Global().OnProgress(func() {
		ok = true
	})
	if !ok {
		t.Error(ok)
	}
}
