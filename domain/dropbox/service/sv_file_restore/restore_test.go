package sv_file_restore

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestRestoreImpl_Restore(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Restore(qtr_endtoend.NewTestDropboxFolderPath(), "test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
