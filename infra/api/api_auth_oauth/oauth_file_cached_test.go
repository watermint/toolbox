package api_auth_oauth_test

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
	"time"
)

func TestCached_Auth(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		name := "test-cached-auth-" + time.Now().String()
		a := api_auth_oauth.NewConsoleCacheOnly(ctl, name, dbx_auth.NewLegacyApp(ctl))
		_, err := a.Start([]string{"test-cached-auth"})
		if err == nil {
			// should not exist
			t.Error("invalid")
		}
	})
}
