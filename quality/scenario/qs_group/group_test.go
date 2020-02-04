package qs_group

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	"strings"
	"testing"
	"time"
)

func TestGroup(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		testGroupName := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		testMemberEmail := ""

		// Create group
		{
			c, err := app_control_impl.Fork(ctl, "group-add")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &group.Add{}, func(r rc_recipe.Recipe) {
				m := r.(*group.Add)
				m.Name = testGroupName
			})
			if err != nil {
				t.Error(err)
				return
			}
		}

		// Delete group
		defer func() {
			c, err := app_control_impl.Fork(ctl, "group-delete")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &group.Delete{}, func(r rc_recipe.Recipe) {
				m := r.(*group.Delete)
				m.Name = testGroupName
			})
			if err != nil {
				t.Error(err)
			}
		}()

		// Verify created group
		{
			c, err := app_control_impl.Fork(ctl, "group-list")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &group.List{}, func(r rc_recipe.Recipe) {})
			if err != nil {
				t.Error(err)
				return
			}

			found := false
			err = qt_recipe.TestRows(c, "group", func(cols map[string]string) error {
				if cols["group_name"] == testGroupName {
					found = true
				}
				return nil
			})
			if err != nil {
				t.Error(err)
			}
			if !found {
				t.Error("test group not found in the list")
			}
		}

		// Member list for adding member to the group
		{
			c, err := app_control_impl.Fork(ctl, "member-list")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &member.List{}, func(r rc_recipe.Recipe) {})
			if err != nil {
				t.Error(err)
				return
			}

			err = qt_recipe.TestRows(c, "member", func(cols map[string]string) error {
				if strings.HasSuffix(cols["email"], "#") {
					return nil
				}
				if cols["status"] != "active" {
					return nil
				}
				testMemberEmail = cols["email"]
				return nil
			})
			if err != nil {
				t.Error(err)
			}
		}

		// Add a member to the group
		{
			c, err := app_control_impl.Fork(ctl, "group-member-add")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &groupmember.Add{}, func(r rc_recipe.Recipe) {
				m := r.(*groupmember.Add)
				m.GroupName = testGroupName
				m.MemberEmail = testMemberEmail
			})
			if err != nil {
				t.Error(err)
				return
			}
		}

		// Verify a member of the group
		{
			c, err := app_control_impl.Fork(ctl, "group-member-list")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &groupmember.List{}, func(r rc_recipe.Recipe) {})
			if err != nil {
				t.Error(err)
				return
			}

			foundGroup := false
			foundMember := false
			err = qt_recipe.TestRows(c, "group_member", func(cols map[string]string) error {
				if cols["group_name"] == testGroupName {
					foundGroup = true
					if cols["email"] == testMemberEmail {
						foundMember = true
					}
				}
				return nil
			})
			if err != nil {
				t.Error(err)
			}
			if !foundGroup {
				t.Error("test group not found in the list")
			}
			if !foundMember {
				t.Error("test group member not found")
			}
		}

		// Remove a member to the group
		{
			c, err := app_control_impl.Fork(ctl, "group-member-delete")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &groupmember.Delete{}, func(r rc_recipe.Recipe) {
				m := r.(*groupmember.Delete)
				m.GroupName = testGroupName
				m.MemberEmail = testMemberEmail
			})
			if err != nil {
				t.Error(err)
				return
			}
		}

		// Verify a member of the group
		{
			c, err := app_control_impl.Fork(ctl, "group-member-list-after-delete")
			if err != nil {
				t.Error(err)
				return
			}
			err = rc_exec.Exec(c, &groupmember.List{}, func(r rc_recipe.Recipe) {})
			if err != nil {
				t.Error(err)
				return
			}

			foundMember := false
			err = qt_recipe.TestRows(c, "group_member", func(cols map[string]string) error {
				if cols["group_name"] == testGroupName {
					if cols["email"] == testMemberEmail {
						foundMember = true
					}
				}
				return nil
			})
			if err != nil {
				t.Error(err)
			}
			if foundMember {
				t.Error("test group member found")
			}
		}
	})
}
