package sv_teamfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEndToEndTeamFolderImpl_List(t *testing.T) {
	qt_api.DoTestBusinessFile(func(ctx api_context.Context) {
		svc := New(ctx)
		list, err := svc.List()
		if err != nil {
			t.Error(err)
			return
		}

		for _, tf := range list {
			if tf.TeamFolderId == "" {
				t.Error("invalid")
			}
			r, err := svc.Resolve(tf.TeamFolderId)
			if err != nil {
				t.Error(err)
			}
			if r.TeamFolderId != tf.TeamFolderId || r.Name != tf.Name {
				t.Error("invalid")
			}
		}
	})
}

// Mock test

func TestTeamFolderImpl_Activate(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Activate(&mo_teamfolder.TeamFolder{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamFolderImpl_Archive(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Archive(&mo_teamfolder.TeamFolder{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamFolderImpl_Create(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Create("test", SyncDefault())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		_, err = sv.Create("test", SyncNoSync())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamFolderImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamFolderImpl_PermDelete(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.PermDelete(&mo_teamfolder.TeamFolder{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamFolderImpl_Rename(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Rename(&mo_teamfolder.TeamFolder{}, "test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamFolderImpl_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
