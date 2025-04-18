package uc_compare_local

import (
	"os"
	"testing"

	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestCompareImpl_Diff(t *testing.T) {
	d, err := os.MkdirTemp("", "compare")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(d)
	}()

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_client_impl.NewMock("mock", ctl)
		uc := New(ctx, ctl.UI())
		_, err := uc.Diff(mo_path2.NewFileSystemPath(d), qtr_endtoend.NewTestDropboxFolderPath(), func(diff mo_file_diff.Diff) error {
			return nil
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
