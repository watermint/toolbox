package qs_group

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	"strings"
	"testing"
	"time"
)

func TestGroup(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		testGroupName := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		testMemberEmail := ""

		// Create group
		qtr_endtoend.ForkWithName(t, "group-add", ctl, func(c app_control.Control) error {
			return rc_exec.Exec(c, &group.Add{}, func(r rc_recipe.Recipe) {
				m := r.(*group.Add)
				m.Name = testGroupName
			})
		})

		// Delete group
		defer qtr_endtoend.ForkWithName(t, "group-delete", ctl, func(c app_control.Control) error {
			return rc_exec.Exec(c, &group.Delete{}, func(r rc_recipe.Recipe) {
				m := r.(*group.Delete)
				m.Name = testGroupName
			})
		})

		// Rename group
		qtr_endtoend.ForkWithName(t, "group-rename", ctl, func(c app_control.Control) error {
			return rc_exec.Exec(c, &group.Rename{}, func(r rc_recipe.Recipe) {
				m := r.(*group.Rename)
				m.CurrentName = testGroupName
				m.NewName = testGroupName + "New"
			})
		})

		// Revert: Rename group
		qtr_endtoend.ForkWithName(t, "group-rename-revert", ctl, func(c app_control.Control) error {
			return rc_exec.Exec(c, &group.Rename{}, func(r rc_recipe.Recipe) {
				m := r.(*group.Rename)
				m.CurrentName = testGroupName + "New"
				m.NewName = testGroupName
			})
		})

		// Verify created group
		qtr_endtoend.ForkWithName(t, "group-rename-list", ctl, func(c app_control.Control) error {
			err, cnt := qt_errors.ErrorsForTest(c.Log(), rc_exec.Exec(c, &group.List{}, func(r rc_recipe.Recipe) {}))
			if !cnt {
				return nil
			}
			if err != nil {
				t.Error(err)
				return err
			}
			found := false
			err = qtr_endtoend.TestRows(c, "group", func(cols map[string]string) error {
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
			return nil
		})

		// Member list for adding member to the group
		qtr_endtoend.ForkWithName(t, "member-list", ctl, func(c app_control.Control) error {
			err, cnt := qt_errors.ErrorsForTest(ctl.Log(), rc_exec.Exec(c, &member.List{}, func(r rc_recipe.Recipe) {}))
			if !cnt {
				return nil
			}
			if err != nil {
				t.Error(err)
				return err
			}

			err = qtr_endtoend.TestRows(c, "member", func(cols map[string]string) error {
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
			return nil
		})

		// Add a member to the group
		qtr_endtoend.ForkWithName(t, "group-member-add", ctl, func(c app_control.Control) error {
			return rc_exec.Exec(c, &groupmember.Add{}, func(r rc_recipe.Recipe) {
				m := r.(*groupmember.Add)
				m.GroupName = testGroupName
				if testMemberEmail == "" {
					m.MemberEmail = "john@example.com"
				} else {
					m.MemberEmail = testMemberEmail
				}
			})
		})

		// Verify a member of the group
		qtr_endtoend.ForkWithName(t, "group-member-list", ctl, func(c app_control.Control) error {
			err, cnt := qt_errors.ErrorsForTest(ctl.Log(), rc_exec.Exec(c, &groupmember.List{}, func(r rc_recipe.Recipe) {}))
			if !cnt {
				return nil
			}
			if err != nil {
				t.Error(err)
				return err
			}

			foundGroup := false
			foundMember := false
			err = qtr_endtoend.TestRows(c, "group_member", func(cols map[string]string) error {
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
			return nil
		})

		// Remove a member to the group
		qtr_endtoend.ForkWithName(t, "group-member-delete", ctl, func(c app_control.Control) error {
			return rc_exec.Exec(c, &groupmember.Delete{}, func(r rc_recipe.Recipe) {
				m := r.(*groupmember.Delete)
				m.GroupName = testGroupName
				if testMemberEmail == "" {
					m.MemberEmail = "john@example.com"
				} else {
					m.MemberEmail = testMemberEmail
				}
			})
		})

		// Verify a member of the group
		qtr_endtoend.ForkWithName(t, "group-member-list-after-delete", ctl, func(c app_control.Control) error {
			err, cnt := qt_errors.ErrorsForTest(ctl.Log(), rc_exec.Exec(c, &groupmember.List{}, func(r rc_recipe.Recipe) {}))
			if !cnt {
				return nil
			}
			if err != nil {
				t.Error(err)
				return nil
			}

			foundMember := false
			err = qtr_endtoend.TestRows(c, "group_member", func(cols map[string]string) error {
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
			return nil
		})
	})
}
