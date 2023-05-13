package legalhold

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_legalhold"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type Release struct {
	Peer     dbx_conn.ConnScopedTeam
	PolicyId string
}

func (z *Release) Preset() {
	//z.Peer.SetScopes(
	//dbx_auth.ScopeMembersRead,
	//dbx_auth.ScopeTeamDataGovernanceWrite,
	//)
}

func (z *Release) Exec(c app_control.Control) error {
	return sv_legalhold.New(z.Peer.Client()).Release(z.PolicyId)
}

func (z *Release) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Release{}, func(r rc_recipe.Recipe) {
		m := r.(*Release)
		m.PolicyId = "pid_dbhid:xxxxx"
	})
}
