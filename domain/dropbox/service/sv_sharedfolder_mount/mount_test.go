package sv_sharedfolder_mount

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEndToEndMountImpl_List(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		mounts, err := svc.List()
		if err != nil {
			t.Error(err)
			return
		}

		for _, m := range mounts {
			if m.SharedFolderId == "" || m.Name == "" {
				t.Error("invalid")
			}
		}
	})
}

// mock test

func TestMountImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMountImpl_Mount(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Mount(&mo_sharedfolder.SharedFolder{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMountImpl_Unmount(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.Unmount(&mo_sharedfolder.SharedFolder{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
