package sv_sharedfolder

import (
	"fmt"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"go.uber.org/zap"
	"strings"
	"testing"
	"time"
)

func TestSharedFolderImpl_List(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
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

func TestSharedFolderImpl_Create(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		svc := New(ctx)
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		ctx.Log().Info("create shared folder", zap.String("name", name))
		sf, err := svc.Create(api_test.ToolboxTestSuiteFolder.ChildPath(name))
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
			if strings.HasPrefix(api_util.ErrorSummary(err), "internal_error") {
				ctx.Log().Warn("Internal error. Ignored")
			} else {
				t.Error("invalid", err)
			}
		}
	})
}
