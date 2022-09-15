package role

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_user"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Add struct {
	Peer   dbx_conn.ConnScopedTeam
	Group  string
	RoleId string
	Roles  rp_model.RowReport
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
	z.Roles.SetModel(
		&mo_adminrole.MemberRole{},
	)
}

func (z *Add) reportRoles(member *mo_member.Member, roles []*mo_adminrole.Role) {
	for _, role := range roles {
		z.Roles.Row(mo_adminrole.NewMemberRole(member.TeamMemberId, member.Email, role))
	}
}

func (z *Add) addRoleMember(member *mo_member.Member, c app_control.Control) error {
	roleIds := member.RoleIds()
	for _, roleId := range roleIds {
		if roleId == z.RoleId {
			z.reportRoles(member, member.Roles())
			return nil
		}
	}

	roleIds = append(roleIds, z.RoleId)

	updated, err := sv_adminrole.New(z.Peer.Client()).UpdateRole(mo_user.NewUserSelectorByTeamMemberId(member.TeamMemberId), roleIds)
	if err != nil {
		return err
	}

	z.reportRoles(member, updated)
	return nil
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.Roles.Open(); err != nil {
		return err
	}

	group, err := sv_group.New(z.Peer.Client()).ResolveByName(z.Group)
	if err != nil {
		return err
	}
	members, err := sv_group_member.New(z.Peer.Client(), group).List()
	if err != nil {
		return err
	}

	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("add_role", z.addRoleMember, c)
		q := s.Get("add_role")

		for _, member := range members {
			m, err := sv_member.New(z.Peer.Client()).Resolve(member.TeamMemberId)
			if err != nil {
				lastErr = err
				continue
			}
			q.Enqueue(m)
		}
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))
	return lastErr
}

func (z *Add) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Group = "CorpIt"
		m.RoleId = "pid_dbtmr:1234"
	})
}
