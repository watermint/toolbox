package sq_sharedfolder

import (
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/sequence/sq_test"
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/infra/api/api_test"
	"github.com/watermint/toolbox/infra/api/api_util"
	"go.uber.org/zap"
	"strings"
	"testing"
	"time"
)

func TestAddUser_Do(t *testing.T) {
	sq_test.DoTestTeamTask(func(biz service.Business) {
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		l := biz.Log().With(zap.String("name", name))

		folderOwner := biz.Admin()

		l.Info("Prepare member list")
		members, err := biz.Member().List()
		if err != nil {
			t.Error("unable to list members", err)
			return
		}

		var coworker *mo_member.Member = nil
		for _, m := range members {
			if m.EmailVerified && m.TeamMemberId != folderOwner.TeamMemberId {
				coworker = m
				break
			}
		}
		if coworker == nil {
			t.Error("No appropriate coworker account found")
			return
		}

		l.Info("Prepare shared folder")
		path := api_test.ToolboxTestSuiteFolder.ChildPath(name)
		sf, err := biz.SharedFolderAsMember(folderOwner.TeamMemberId).Create(path)
		if err != nil {
			t.Error("unable to create shared folder", err)
			return
		}

		task := AddUser{
			SharedFolderId: sf.SharedFolderId,
			UserEmail:      coworker.Email,
			AccessLevel:    mo_sharedfolder_member.AccessTypeEditor,
		}

		l.Info("Do task")
		if err = task.Do(biz); err != nil {
			t.Error("task failed", err)
		}

		l.Info("Verify")
		sfMembers, err := biz.SharedFolderMemberAsMember(sf.SharedFolderId, folderOwner.TeamMemberId).List()
		if err != nil {
			t.Error("unable to list shared folder members", err)
		} else {
			found := false
			for _, sm := range sfMembers {
				if u, e := sm.User(); e {
					if u.TeamMemberId == coworker.TeamMemberId && u.AccessType() == task.AccessLevel {
						found = true
						break
					}
				}
			}
			if !found {
				t.Error("coworker not found, or invalid access type")
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
