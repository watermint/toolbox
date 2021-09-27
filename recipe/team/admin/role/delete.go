package role

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_user"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Delete struct {
	Peer                dbx_conn.ConnScopedTeam
	Email               string
	RoleId              string
	Roles               rp_model.RowReport
	SkipDoesNotHaveRole app_msg.Message
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
	z.Roles.SetModel(
		&mo_adminrole.MemberRole{},
	)
}

func (z *Delete) reportRoles(member *mo_member.Member, roles []*mo_adminrole.Role) {
	for _, role := range roles {
		z.Roles.Row(mo_adminrole.NewMemberRole(member.TeamMemberId, member.Email, role))
	}
}

func (z *Delete) Exec(c app_control.Control) error {
	if err := z.Roles.Open(); err != nil {
		return err
	}

	member, err := sv_member.New(z.Peer.Context()).ResolveByEmail(z.Email)
	if err != nil {
		return err
	}

	newRoleIds := make([]string, 0)
	found := false
	roleIds := member.RoleIds()
	for _, roleId := range roleIds {
		if roleId == z.RoleId {
			found = true
		} else {
			newRoleIds = append(newRoleIds, roleId)
		}
	}

	if !found {
		c.UI().Info(z.SkipDoesNotHaveRole)
		return nil
	}

	roleIds = append(roleIds, z.RoleId)

	updated, err := sv_adminrole.New(z.Peer.Context()).UpdateRole(mo_user.NewUserSelectorByTeamMemberId(member.TeamMemberId), newRoleIds)
	if err != nil {
		return err
	}

	z.reportRoles(member, updated)
	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Email = "jo@example.com"
		m.RoleId = "pid_dbtmr:1234"
	})
}
