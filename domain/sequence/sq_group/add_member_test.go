package sq_group

import (
	"fmt"
	"github.com/watermint/toolbox/domain/sequence/sq_test"
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestAddMember_Do(t *testing.T) {
	sq_test.DoTestTeamTask(func(biz service.Business) {
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		l := biz.Log().With(zap.String("name", name))

		l.Info("Prepare")
		group, err := biz.Group().Create(name, sv_group.CompanyManaged())
		if err != nil {
			t.Error("unable to create group", err)
			return
		}

		targetMember := biz.Admin()

		task := AddMember{
			GroupName:   name,
			MemberEmail: targetMember.Email,
		}

		l.Info("Do task")
		if err = task.Do(biz); err != nil {
			t.Error("task failed", err)
		}

		l.Info("Verify")
		groupMembers, err := biz.GroupMember(group.GroupId).List()
		if err != nil {
			t.Error("unable to list group member", err)
		} else {
			found := false
			for _, gm := range groupMembers {
				if gm.TeamMemberId == targetMember.TeamMemberId {
					found = true
					break
				}
			}
			if !found {
				t.Error("member not found")
			}
		}

		l.Info("Clean up")
		err = biz.Group().Remove(group.GroupId)
		if err != nil {
			t.Error("unable to clean up", err)
		}
	})
}
