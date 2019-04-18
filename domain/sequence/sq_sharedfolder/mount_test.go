package sq_sharedfolder

import (
	"fmt"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/sequence"
	"github.com/watermint/toolbox/domain/service"
	"go.uber.org/zap"
	"strings"
	"testing"
	"time"
)

func TestMount_Do(t *testing.T) {
	sequence.DoTestTeamTask(func(biz service.Business) {
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		l := biz.Log().With(zap.String("name", name))

		folderOwner := biz.Admin()

		l.Info("Prepare shared folder")
		path := api_test.ToolboxTestSuiteFolder.ChildPath(name)
		sf, err := biz.SharedFolderAsMember(folderOwner.TeamMemberId).Create(path)
		if err != nil {
			t.Error("unable to create shared folder", err)
			return
		}

		l.Info("Unmount")
		err = biz.SharedFolderMountAsMember(folderOwner.TeamMemberId).Unmount(sf)
		if err != nil {
			t.Error("unable to unmount", err)
			err = biz.SharedFolderAsMember(folderOwner.TeamMemberId).Remove(sf)
			if err != nil {
				t.Error("unable to clean up", err)
			}
			return
		}

		task := Mount{
			SharedFolderId: sf.SharedFolderId,
			UserEmail:      folderOwner.Email,
			MountPoint:     path.Path(),
		}

		l.Info("Do task")
		if err = task.Do(biz); err != nil {
			t.Error("task failed", err)
		}

		l.Info("Verify")
		sf, err = biz.SharedFolderAsMember(folderOwner.TeamMemberId).Resolve(sf.SharedFolderId)
		if err != nil {
			t.Error("unable to resolve shared folder", err)
		} else {
			l.Debug("Compare path", zap.String("pathLower", sf.PathLower), zap.String("path", path.Path()))
			if sf.PathLower != strings.ToLower(path.Path()) {
				t.Error("path miss match", sf.PathLower, path.Path())
			}
		}

		l.Info("Clean up")

		err = biz.SharedFolderAsMember(folderOwner.TeamMemberId).Remove(sf)
		if err != nil {
			if strings.HasPrefix(api_util.ErrorSummary(err), "internal_error") {
				l.Warn("Internal error. Ignored")
			} else {
				t.Error("unable to clean up", err)
			}
		}
	})
}
