package admin

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type List struct {
	Peer            dbx_conn.ConnScopedTeam
	MemberRoles     da_griddata.GridDataOutput
	IncludeNonAdmin bool
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
	)
}

func (z *List) Exec(c app_control.Control) error {
	roles, err := sv_adminrole.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	header := make([]interface{}, len(roles)+2)
	header[0] = "team_member_id"
	header[1] = "email"
	for i := 0; i < len(roles); i++ {
		header[2+i] = roles[i].Name
	}
	z.MemberRoles.Row(header)

	return sv_member.New(z.Peer.Context()).ListEach(func(member *mo_member.Member) bool {
		row := make([]interface{}, len(roles)+2)
		if !z.IncludeNonAdmin && len(member.RoleIds()) < 1 {
			return true
		}
		row[0] = member.TeamMemberId
		row[1] = member.Email
		for i := 0; i < len(roles); i++ {
			row[2+i] = ""
			for _, mr := range member.Roles() {
				if mr.RoleId == roles[i].RoleId {
					row[2+i] = "true"
					break
				}
			}
		}
		z.MemberRoles.Row(row)
		return true
	})
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
