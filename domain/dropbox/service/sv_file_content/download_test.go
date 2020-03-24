package sv_file_content

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestDownloadImpl_Download(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := NewDownload(ctx)
		_, _, err := sv.Download(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
