package role

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_user"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_adminrole"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Clear struct {
	Peer  dbx_conn.ConnScopedTeam
	Email string
}

func (z *Clear) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
}

func (z *Clear) Exec(c app_control.Control) error {
	member, err := sv_member.New(z.Peer.Context()).ResolveByEmail(z.Email)
	if err != nil {
		return err
	}

	_, err = sv_adminrole.New(z.Peer.Context()).UpdateRole(mo_user.NewUserSelectorByTeamMemberId(member.TeamMemberId), []string{})
	if err != nil {
		return err
	}

	return nil
}

func (z *Clear) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Clear{}, func(r rc_recipe.Recipe) {
		m := r.(*Clear)
		m.Email = "jo@example.com"
	})
}
