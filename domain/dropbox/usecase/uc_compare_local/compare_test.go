package uc_compare_local

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"io/ioutil"
	"os"
	"testing"
)

func TestCompareImpl_Diff(t *testing.T) {
	d, err := ioutil.TempDir("", "compare")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		os.RemoveAll(d)
	}()

	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		ctx := dbx_context.NewMock(ctl)
		uc := New(ctx, ctl.UI())
		_, err := uc.Diff(mo_path.NewFileSystemPath(d), qt_recipe.NewTestDropboxFolderPath(), func(diff mo_file_diff.Diff) error {
			return nil
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
