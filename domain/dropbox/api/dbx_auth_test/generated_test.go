package dbx_auth

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestGenerated_Auth(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		a := dbx_auth.NewConsoleGenerated(ctl, "test-generated-auth")
		if a.PeerName() != "test-generated-auth" {
			t.Error(a.PeerName())
		}
		_, err := a.Auth(api_auth.DropboxTokenBusinessInfo)
		if err != app.ErrorUserCancelled {
			t.Error(err)
		}
	})
}
