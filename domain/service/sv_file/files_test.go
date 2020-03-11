package sv_file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestFilesImpl_ListWithTestSuite(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		svc := newFilesTest(ctx)
		folder := qt_api.ToolboxTestSuiteFolder.ChildPath("list_folder")
		entries, err := svc.List(folder)
		if err != nil {
			t.Error(err)
			return
		}
		if len(entries) < 1 {
			t.Error("invalid")
		}
		for i, e := range entries {
			if i > 10 {
				break
			}
			f, err := svc.Resolve(e.Path())
			if err != nil {
				t.Error(err)
			}
			if f.Tag() != e.Tag() || f.PathLower() != e.PathLower() {
				t.Error("invalid")
			}
		}
	})
}

func TestFilesImpl_ListChunked(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
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
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
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
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewFiles(ctx)
		err := sv.Poll(qt_recipe.NewTestDropboxFolderPath(), func(entry mo_file.Entry) {
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewFiles(ctx)
		_, err := sv.Remove(qt_recipe.NewTestDropboxFolderPath(), RemoveRevision("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewFiles(ctx)
		_, err := sv.Resolve(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFilesImpl_Search(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
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
