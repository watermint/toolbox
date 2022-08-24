package sv_filerequest

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFileRequestImpl_Create(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Create("test", qtr_endtoend.NewTestDropboxFolderPath(),
			OptDeadline("2020-03-02T17:40:00Z"),
			OptAllowLateUploads(""))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFileRequestImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFileRequestImpl_Delete(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Delete("1234")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFileRequestImpl_DeleteAllClosed(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.DeleteAllClosed()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFileRequestImpl_Update(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Update(&mo_filerequest.FileRequest{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
