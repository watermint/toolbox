package sv_sharedfolder

import (
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestSharedFolderImpl_List(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		folders, err := svc.List()
		if err != nil {
			t.Error(err)
			return
		}
		for _, f := range folders {
			if f.Name == "" {
				t.Error("invalid")
			}
			if f.SharedFolderId == "" {
				t.Error("invalid")
			}
		}
	})
}

func TestEndToEndSharedFolderImpl_Create(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		ctx.Log().Info("create shared folder", zap.String("name", name))
		sf, err := svc.Create(qt_api.ToolboxTestSuiteFolder.ChildPath(name))
		if err != nil {
			t.Error("invalid", err)
		}
		if sf.Name != name {
			t.Error("invalid")
		}
		ctx.Log().Info("create shared folder again", zap.String("name", name))

		// Response on already_shared
		// {
		//  "error_summary": "bad_path/already_shared/...",
		//  "error": {
		//    ".tag": "bad_path",
		//    "bad_path": {
		//      ".tag": "already_shared",
		//      "access_type": {
		//        ".tag": "owner"
		//      },
		//      "is_inside_team_folder": false,
		//      "is_team_folder": false,
		//      "path_lower": "/xxxxxx-xxx-xxxxxxxxxx",
		//      "name": "xxxxxx-xxx-xxxxxxxxxx",
		//      "policy": {
		//        "acl_update_policy": {
		//          ".tag": "editors"
		//        },
		//        "shared_link_policy": {
		//          ".tag": "anyone"
		//        },
		//        "viewer_info_policy": {
		//          ".tag": "enabled"
		//        }
		//      },
		//      "preview_url": "https://www.dropbox.com/scl/fo/xxxxxxxxxxxxxxxxxxxxx/xxxxxxxxxxxxxx-xxxxxxxxxx?dl=0",
		//      "shared_folder_id": "xxxxxxxxxx",
		//      "time_invited": "2019-04-01T07:34:44Z",
		//      "access_inheritance": {
		//        ".tag": "inherit"
		//      }
		//    }
		//  },
		//  "user_message": {
		//    "locale": "en",
		//    "text": "This folder is already shared."
		//  }
		// }

		//sf, err = svc.Create(api_test.ToolboxTestSuiteFolder.ChildPath(name))
		//if err != nil {
		//	eb := api_util.ErrorBody(err)
		//	if eb != nil {
		//		ctx.Log().Error("err", zap.Error(err), zap.String("errorBody", string(eb)))
		//	}
		//	t.Error("invalid", err)
		//}

		err = svc.Remove(sf)
		if err != nil {
			t.Error("invalid", err)
		}
	})
}

// Mock tests

func TestSharedFolderImpl_Create(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Create(qt_recipe.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Leave(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.Leave(&mo_sharedfolder.SharedFolder{}, LeaveACopy())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.Remove(&mo_sharedfolder.SharedFolder{}, LeaveACopy())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Transfer(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.Transfer(&mo_sharedfolder.SharedFolder{}, ToProfile(&mo_profile.Profile{}))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Transfer(&mo_sharedfolder.SharedFolder{}, ToAccountId("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Transfer(&mo_sharedfolder.SharedFolder{}, ToTeamMemberId("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_UpdatePolicy(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.UpdatePolicy("test", MemberPolicy("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
