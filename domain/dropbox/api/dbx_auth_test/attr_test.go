package dbx_auth

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth_attr"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestAttr_Auth(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ma := dbx_auth.NewMock("test-mock")
		aa := dbx_auth_attr.NewConsoleAttr(ctl, ma)
		if aa.PeerName() != "test-mock" {
			t.Error(aa.PeerName())
		}
		if _, err := aa.Auth([]string{"test-scope"}); err != nil {
			t.Error(err)
		}
		if _, err := aa.Auth([]string{api_auth.DropboxTokenFull}); err == nil {
			t.Error("invalid")
		}
		if _, err := aa.Auth([]string{api_auth.DropboxTokenBusinessAudit}); err == nil {
			t.Error("invalid")
		}
	})
}
