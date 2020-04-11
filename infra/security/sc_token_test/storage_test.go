package sc_token_test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
	"github.com/watermint/toolbox/infra/security/sc_token"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"golang.org/x/oauth2"
	"testing"
	"time"
)

func TestObfuscatedStorage_PutGet(t *testing.T) {
	qt_recipe.TestWithControl(t, func(c app_control.Control) {
		s := sc_token.NewObfuscated(c, "TestObfuscatedStorage_PutGet")
		s.Purge("test1")

		if _, err := s.Get("test1"); err != sc_obfuscate.ErrorStorageNotFound {
			t.Error(err)
		}
		token0 := &oauth2.Token{
			AccessToken:  "test-access-token",
			TokenType:    "Bearer",
			RefreshToken: "test-refresh-token",
			Expiry:       time.Time{},
		}
		if err := s.Put("test1", token0); err != nil {
			t.Error(err)
		}

		token1, err := s.Get("test1")
		if err != nil {
			t.Error(err)
		}

		if token0.AccessToken != token1.AccessToken || token0.RefreshToken != token1.RefreshToken {
			t.Error("invalid")
		}
	})
}