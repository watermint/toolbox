package sv_file_content

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestExportImpl_Export(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewExport(ctx)
		_, _, err := sv.Export(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
