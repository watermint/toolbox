package dbx_auth

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"testing"
)

func TestNoAuth_Auth(t *testing.T) {
	a := api_auth.NewNoAuth()
	if a.PeerName() != "" || a.Scope() != "" || a.Supplemental() != "" || a.Description() != "" {
		t.Error("invalid")
	}
	if !a.IsNoAuth() {
		t.Error("invalid")
	}
}
