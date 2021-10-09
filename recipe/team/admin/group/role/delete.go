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
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Delete struct {
	Peer           dbx_conn.ConnScopedTeam
	ExceptionGroup string
	RoleId         string
	Roles          rp_model.RowReport
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
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

func (z *Delete) deleteRoleMember(member *mo_member.Member, c app_control.Control) error {
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
		return nil
	}

	updated, err := sv_adminrole.New(z.Peer.Context()).UpdateRole(mo_user.NewUserSelectorByTeamMemberId(member.TeamMemberId), newRoleIds)
	if err != nil {
		return err
	}

	z.reportRoles(member, updated)
	return nil
}

func (z *Delete) Exec(c app_control.Control) error {
	if err := z.Roles.Open(); err != nil {
		return err
	}

	group, err := sv_group.New(z.Peer.Context()).ResolveByName(z.ExceptionGroup)
	if err != nil {
		return err
	}
	exceptionMembers, err := sv_group_member.New(z.Peer.Context(), group).List()
	if err != nil {
		return err
	}

	isTargetMember := func(m *mo_member.Member) bool {
		for _, em := range exceptionMembers {
			if em.TeamMemberId == m.TeamMemberId {
				return false
			}
		}
		return true
	}

	var lastErr, listErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("delete_role", z.deleteRoleMember, c)
		q := s.Get("delete_role")

		listErr = sv_member.New(z.Peer.Context()).ListEach(func(member *mo_member.Member) bool {
			if isTargetMember(member) {
				q.Enqueue(member)
			}
			return true
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))
	return lang.NewMultiErrorOrNull(lastErr, listErr)
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.ExceptionGroup = "CorpIt"
		m.RoleId = "pid_dbtmr:1234"
	})
}
