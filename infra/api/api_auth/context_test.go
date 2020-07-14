package api_auth

import (
	"golang.org/x/oauth2"
	"testing"
)

func TestNoAuth(t *testing.T) {
	a := NewNoAuth()
	if !a.IsNoAuth() {
		t.Error("invalid")
	}
	if a.PeerName() != "" || a.Token() == nil || a.PeerName() != "" ||
		a.Description() != "" || a.Supplemental() != "" || len(a.Scopes()) > 0 {
		t.Error("invalid")
	}
}

func TestContext(t *testing.T) {
	c := NewContext(&oauth2.Token{}, "test-context", []string{"test-scope"})
	if c.IsNoAuth() || c.PeerName() != "test-context" || c.Scopes()[0] != "test-scope" ||
		c.Description() != "" || c.Supplemental() != "" || c.Token() == nil {
		t.Error("invalid")
	}

	c = NewContextWithAttr(c, "test-desc", "test-suppl")
	if c.IsNoAuth() || c.PeerName() != "test-context" || c.Scopes()[0] != "test-scope" ||
		c.Description() != "test-desc" || c.Supplemental() != "test-suppl" || c.Token() == nil {
		t.Error("invalid")
	}
}
