package uc_file_merge

import (
	"github.com/watermint/toolbox/infra/api/api_context_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestMergeImpl_Merge(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		ctx := api_context_impl.NewMock(ctl)
		sv := New(ctx, ctl)
		err := sv.Merge(
			qt_recipe.NewTestDropboxFolderPath("from"),
			qt_recipe.NewTestDropboxFolderPath("to"),
			DryRun(),
			WithinSameNamespace(),
			ClearEmptyFolder(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
