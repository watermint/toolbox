package sq_sharedfolder

import (
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/sequence/sq_test"
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/api/api_test"
	"github.com/watermint/toolbox/infra/api/api_util"
	"go.uber.org/zap"
	"strings"
	"testing"
	"time"
)

func TestAddGroup_Do(t *testing.T) {
	sq_test.DoTestTeamTask(func(biz service.Business) {
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		l := biz.Log().With(zap.String("name", name))

		targetMember := biz.Admin()

		l.Info("Prepare shared folder")
		path := api_test.ToolboxTestSuiteFolder.ChildPath(name)
		sf, err := biz.SharedFolderAsMember(targetMember.TeamMemberId).Create(path)
		if err != nil {
			t.Error("unable to create shared folder", err)
			return
		}

		l.Info("Prepare group")
		group, err := biz.Group().Create(name, sv_group.CompanyManaged())
		if err != nil {
			t.Error("unable to create group", err)
			return
		}

		task := AddGroup{
			SharedFolderId: sf.SharedFolderId,
			GroupName:      group.GroupName,
			AccessLevel:    mo_sharedfolder_member.AccessTypeEditor,
		}

		l.Info("Do task")
		if err = task.Do(biz); err != nil {
			t.Error("task failed", err)
		}

		l.Info("Verify")
		sfMembers, err := biz.SharedFolderMemberAsMember(sf.SharedFolderId, targetMember.TeamMemberId).List()
		if err != nil {
			t.Error("unable to list shared folder members", err)
		} else {
			found := false
			for _, sm := range sfMembers {
				if g, e := sm.Group(); e {
					if g.GroupId == group.GroupId && g.AccessType() == task.AccessLevel {
						found = true
						break
					}
				}
			}
			if !found {
				t.Error("group not found, or invalid access type")
			}
		}

		l.Info("Clean up")
		err = biz.Group().Remove(group.GroupId)
		if err != nil {
			if strings.HasPrefix(api_util.ErrorSummary(err), "internal_error") {
				l.Warn("Internal error. Ignored")
			} else {
				t.Error("unable to clean up", err)
			}
		}

		err = biz.SharedFolderAsMember(targetMember.TeamMemberId).Remove(sf)
		if err != nil {
			if strings.HasPrefix(api_util.ErrorSummary(err), "internal_error") {
				l.Warn("Internal error. Ignored")
			} else {
				t.Error("unable to clean up", err)
			}
		}
	})
}
