package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUploadImpl_Add(t *testing.T) {
	f, err := qt_file.MakeDummyFile("add")
	if err != nil {
		t.Error(err)
		return
	}
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewUpload(ctx)
		_, err := sv.Add(qtr_endtoend.NewTestDropboxFolderPath(), f)
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
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewUpload(ctx)
		_, err := sv.Overwrite(qtr_endtoend.NewTestDropboxFolderPath(), f)
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
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewUpload(ctx)
		_, err := sv.Update(qtr_endtoend.NewTestDropboxFolderPath(), f, "test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
