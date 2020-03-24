package sv_file_revision

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestRevisionImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		_, err := sv.List(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestRevisionImpl_ListById(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		_, err := sv.ListById(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
