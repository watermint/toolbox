package sv_file_restore

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestRestoreImpl_Restore(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Restore(qt_recipe.NewTestDropboxFolderPath(), "test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}