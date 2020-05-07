package uc_file_merge

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestMergeImpl_Merge(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context_impl.NewMock(ctl)
		sv := New(ctx, ctl)
		err := sv.Merge(
			qtr_endtoend.NewTestDropboxFolderPath("from"),
			qtr_endtoend.NewTestDropboxFolderPath("to"),
			DryRun(),
			WithinSameNamespace(),
			ClearEmptyFolder(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
