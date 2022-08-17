package dbx_auth

import (
	"golang.org/x/oauth2"
	"testing"
)

func TestMockConsoleAuth_Auth(t *testing.T) {
	{
		a := NewMock("test-mock")
		if a.PeerName() != "test-mock" {
			t.Error(a.PeerName())
		}
		c, err := a.Start([]string{"test-scope"})
		if err != nil {
			t.Error(err)
		}
		if c.PeerName() != "test-mock" {
			t.Error(c.PeerName())
		}
	}

	{
		p := make(map[string]*oauth2.Token)
		p["test-scope"] = &oauth2.Token{}
		a := NewMockWithPreset("test-mock", p)
		if a.PeerName() != "test-mock" {
			t.Error(a.PeerName())
		}
		c, err := a.Start([]string{"test-scope"})
		if err != nil {
			t.Error(err)
		}
		if c.PeerName() != "test-mock" {
			t.Error(c.PeerName())
		}
	}
}

func TestMockContext(t *testing.T) {
	c := &MockContext{
		peerName: "test-peer",
		scopes:   []string{"test-scope"},
		preset:   &oauth2.Token{},
	}
	if c.IsNoAuth() {
		t.Error("invalid")
	}
	if c.PeerName() != "test-peer" {
		t.Error("invalid")
	}
	if c.Scopes()[0] != "test-scope" {
		t.Error("invalid")
	}
	if c.Description() != "" {
		t.Error("invalid")
	}
	if c.Supplemental() != "" {
		t.Error("invalid")
	}
	if c.Token() == nil {
		t.Error("invalid")
	}
}
