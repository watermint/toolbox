package dbx_auth

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth_attr"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestAttr_Auth(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		ma := dbx_auth.NewMock("test-mock")
		aa := dbx_auth_attr.NewConsoleAttr(ctl, ma)
		if aa.PeerName() != "test-mock" {
			t.Error(aa.PeerName())
		}
		if _, err := aa.Auth("test-scope"); err != dbx_auth_attr.ErrorNoVerification {
			t.Error(err)
		}
		if _, err := aa.Auth(api_auth.DropboxTokenFull); err == nil {
			t.Error("invalid")
		}
		if _, err := aa.Auth(api_auth.DropboxTokenBusinessAudit); err == nil {
			t.Error("invalid")
		}
	})
}
