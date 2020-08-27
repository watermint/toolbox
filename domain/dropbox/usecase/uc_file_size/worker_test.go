package uc_file_size

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestScaleImpl_Size(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context_impl.NewMock(ctl)
		sv := New(ctx, ctl)
		_, errs := sv.Size(qtr_endtoend.NewTestDropboxFolderPath(), 1)
		if len(errs) > 0 {
			for _, e := range errs {
				if e != qt_errors.ErrorMock {
					t.Error(e)
				}
			}
		}
	})
}
