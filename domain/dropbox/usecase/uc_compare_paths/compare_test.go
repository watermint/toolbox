package uc_compare_paths

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

// Mock test

func TestCompareImpl_Diff(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context_impl.NewMock("mock", ctl)
		sv := New(ctx, ctx, ctl.UI())
		_, err := sv.Diff(qtr_endtoend.NewTestDropboxFolderPath(), qtr_endtoend.NewTestDropboxFolderPath(), func(diff mo_file_diff.Diff) error {
			return nil
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
