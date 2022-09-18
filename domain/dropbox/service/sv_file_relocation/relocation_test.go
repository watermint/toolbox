package sv_file_relocation

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestImplRelocation_Copy(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx, AllowOwnershipTransfer(true), AllowSharedFolder(true), AutoRename(true))
		_, err := sv.Copy(qtr_endtoend.NewTestDropboxFolderPath("from"), qtr_endtoend.NewTestDropboxFolderPath("to"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestImplRelocation_Move(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx, AllowOwnershipTransfer(true), AllowSharedFolder(true), AutoRename(true))
		_, err := sv.Move(qtr_endtoend.NewTestDropboxFolderPath("from"), qtr_endtoend.NewTestDropboxFolderPath("to"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
