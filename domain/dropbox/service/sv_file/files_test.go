package sv_file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestFilesImpl_ListChunked(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewFiles(ctx)
		err := sv.ListChunked(qt_recipe.NewTestDropboxFolderPath(), func(entry mo_file.Entry) {},
			Recursive(),
			IncludeMediaInfo(),
			IncludeDeleted(),
			IncludeHasExplicitSharedMembers(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_List(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewFiles(ctx)
		_, err := sv.List(qt_recipe.NewTestDropboxFolderPath(),
			Recursive(),
			IncludeMediaInfo(),
			IncludeDeleted(),
			IncludeHasExplicitSharedMembers(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Poll(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewFiles(ctx)
		err := sv.Poll(qt_recipe.NewTestDropboxFolderPath(), func(entry mo_file.Entry) {
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Remove(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewFiles(ctx)
		_, err := sv.Remove(qt_recipe.NewTestDropboxFolderPath(), RemoveRevision("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Resolve(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewFiles(ctx)
		_, err := sv.Resolve(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Search(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewFiles(ctx)
		_, err := sv.Search("test",
			SearchPath(qt_recipe.NewTestDropboxFolderPath()),
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
