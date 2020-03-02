package sv_sharedlink_file

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_url"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestFileImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.List(mo_url.NewEmptyUrl(), qt_recipe.NewTestDropboxFolderPath(), func(entry mo_file.Entry) {},
			IncludeDeleted(),
			IncludeHasExplicitSharedMembers(),
			Password("test"),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestFileImpl_ListRecursive(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.ListRecursive(mo_url.NewEmptyUrl(), func(entry mo_file.Entry) {},
			IncludeDeleted(),
			IncludeHasExplicitSharedMembers(),
			Password("test"),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
