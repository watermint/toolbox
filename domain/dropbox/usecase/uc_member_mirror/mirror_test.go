package uc_member_mirror

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMirrorImpl_Mirror(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx, ctx)
		err := sv.Mirror("src@example.com", "to@example.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
