package uc_file_mirror

import (
	"github.com/watermint/toolbox/infra/api/dbx_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestFilesImpl_Mirror(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context.NewMock(ctl)
		sv := New(ctx, ctx)
		err := sv.Mirror(qt_recipe.NewTestDropboxFolderPath("from"), qt_recipe.NewTestDropboxFolderPath("to"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
