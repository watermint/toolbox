package uc_file_size

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestScaleImpl_Size(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context.NewMock(ctl)
		sv := New(ctx, ctl)
		_, errs := sv.Size(qt_recipe.NewTestDropboxFolderPath(), 1)
		if len(errs) > 0 {
			for _, e := range errs {
				if e != qt_errors.ErrorMock {
					t.Error(e)
				}
			}
		}
	})
}
