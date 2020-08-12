package eq_pipe_preserve

import "testing"

func TestNopRestorer_Restore(t *testing.T) {
	r := NopRestorer()
	loader := func(d []byte) error {
		t.Error("should not restore any data")
		return nil
	}
	err := r.Restore(loader, loader)
	if err != nil {
		t.Error(err)
	}
}
