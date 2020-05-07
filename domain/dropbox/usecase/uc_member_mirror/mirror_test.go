package uc_member_mirror

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestMirrorImpl_Mirror(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx, ctx)
		err := sv.Mirror("src@example.com", "to@example.com")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
