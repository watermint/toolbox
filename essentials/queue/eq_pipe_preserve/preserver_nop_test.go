package eq_pipe_preserve

import "testing"

func TestNopPreserver(t *testing.T) {
	p := NopPreserver()
	if x := p.Start(); x != ErrorSessionIsNotAvailable {
		t.Error(x)
	}
	if x := p.Add([]byte("Hello")); x != ErrorSessionIsNotAvailable {
		t.Error(x)
	}
	if y, x := p.Commit([]byte("{}")); x != ErrorSessionIsNotAvailable {
		t.Error(y, x)
	}
}
