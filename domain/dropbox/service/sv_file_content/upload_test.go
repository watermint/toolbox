package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestUploadImpl_Add(t *testing.T) {
	f, err := qt_file.MakeDummyFile("add")
	if err != nil {
		t.Error(err)
		return
	}
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewUpload(ctx)
		_, err := sv.Add(qt_recipe.NewTestDropboxFolderPath(), f)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestUploadImpl_Overwrite(t *testing.T) {
	f, err := qt_file.MakeDummyFile("overwrite")
	if err != nil {
		t.Error(err)
		return
	}
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewUpload(ctx)
		_, err := sv.Overwrite(qt_recipe.NewTestDropboxFolderPath(), f)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestUploadImpl_Update(t *testing.T) {
	f, err := qt_file.MakeDummyFile("update")
	if err != nil {
		t.Error(err)
		return
	}
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewUpload(ctx)
		_, err := sv.Update(qt_recipe.NewTestDropboxFolderPath(), f, "test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
