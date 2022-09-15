package sv_file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFilesImpl_ListChunked(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewFiles(ctx)
		err := sv.ListEach(qtr_endtoend.NewTestDropboxFolderPath(), func(entry mo_file.Entry) {},
			Recursive(true),
			IncludeDeleted(true),
			IncludeHasExplicitSharedMembers(true),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewFiles(ctx)
		_, err := sv.List(qtr_endtoend.NewTestDropboxFolderPath(),
			Recursive(true),
			IncludeDeleted(true),
			IncludeHasExplicitSharedMembers(true),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Poll(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewFiles(ctx)
		err := sv.Poll(qtr_endtoend.NewTestDropboxFolderPath(), func(entry mo_file.Entry) {
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Remove(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewFiles(ctx)
		_, err := sv.Remove(qtr_endtoend.NewTestDropboxFolderPath(), RemoveRevision("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewFiles(ctx)
		_, err := sv.Resolve(qtr_endtoend.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Search(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := NewFiles(ctx)
		_, err := sv.Search("test",
			SearchPath(qtr_endtoend.NewTestDropboxFolderPath()),
			SearchMaxResults(100),
			SearchFileDeleted(),
			SearchFileNameOnly(),
			SearchFileExtension("test"),
			SearchCategories("pdf"),
			SearchIncludeHighlights(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
