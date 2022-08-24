package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestExportImpl_Export(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewExport(ctx)
		_, _, err := sv.Export(qtr_endtoend.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
