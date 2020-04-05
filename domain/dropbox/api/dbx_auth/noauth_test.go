package dbx_auth

import "testing"

func TestNoAuth_Auth(t *testing.T) {
	a := NewNoAuth()
	if a.PeerName() != "" || a.Scope() != "" || a.Supplemental() != "" || a.Description() != "" {
		t.Error("invalid")
	}
	if !a.IsNoAuth() {
		t.Error("invalid")
	}
}
