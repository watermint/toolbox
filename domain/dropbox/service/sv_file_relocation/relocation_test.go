package sv_file_relocation

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestImplRelocation_Copy(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx, AllowOwnershipTransfer(true), AllowSharedFolder(true), AutoRename(true))
		_, err := sv.Copy(qt_recipe.NewTestDropboxFolderPath("from"), qt_recipe.NewTestDropboxFolderPath("to"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestImplRelocation_Move(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx, AllowOwnershipTransfer(true), AllowSharedFolder(true), AutoRename(true))
		_, err := sv.Move(qt_recipe.NewTestDropboxFolderPath("from"), qt_recipe.NewTestDropboxFolderPath("to"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
