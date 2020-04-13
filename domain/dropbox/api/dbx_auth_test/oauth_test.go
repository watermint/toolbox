package dbx_auth_test

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestOAuth_Auth(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		a := api_auth_impl.NewConsoleOAuth(ctl, "test-oauth-auth")
		if a.PeerName() != "test-oauth-auth" {
			t.Error(a.PeerName())
		}
		_, err := a.Auth("test-scope")
		if err != api_auth.ErrorUserCancelled {
			t.Error(err)
		}
	})
}
