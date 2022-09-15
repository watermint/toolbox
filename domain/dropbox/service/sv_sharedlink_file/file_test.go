package sv_sharedlink_file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestFileImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.List(mo_url.NewEmptyUrl(), qtr_endtoend.NewTestDropboxFolderPath(), func(entry mo_file.Entry) {},
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
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
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
