package sv_file_relocation

import (
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		r := New(ctx)
		src := qt_api.ToolboxTestSuiteFolder.ChildPath("copy/F0.jpg")
		name := fmt.Sprintf("copy-%x.jpg", time.Now().Unix())
		dest := qt_api.ToolboxTestSuiteFolder.ChildPath("copy", name)

		entry, err := r.Copy(src, dest)
		if err != nil {
			t.Error(err)
		}
		if entry.Name() != name {
			t.Error("invalid")
		}
	})
}

func TestImplRelocation_Copy(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx, AllowOwnershipTransfer(true), AllowSharedFolder(true), AutoRename(true))
		_, err := sv.Copy(qt_recipe.NewTestDropboxFolderPath("from"), qt_recipe.NewTestDropboxFolderPath("to"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestImplRelocation_Move(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx, AllowOwnershipTransfer(true), AllowSharedFolder(true), AutoRename(true))
		_, err := sv.Move(qt_recipe.NewTestDropboxFolderPath("from"), qt_recipe.NewTestDropboxFolderPath("to"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
