package update

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_legalhold"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_legalhold"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Name struct {
	Peer     dbx_conn.ConnScopedTeam
	PolicyId string
	Name     string
	Policy   rp_model.RowReport
}

func (z *Name) Preset() {
	//z.Peer.SetScopes(
	//dbx_auth.ScopeMembersRead,
	//dbx_auth.ScopeTeamDataGovernanceWrite,
	//)
	z.Policy.SetModel(&mo_legalhold.Policy{})
}

func (z *Name) Exec(c app_control.Control) error {
	if err := z.Policy.Open(); err != nil {
		return err
	}

	policy, err := sv_legalhold.New(z.Peer.Client()).UpdateName(
		z.PolicyId,
		z.Name,
	)
	if err != nil {
		return err
	}

	z.Policy.Row(policy)
	return nil
}

func (z *Name) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Name{}, func(r rc_recipe.Recipe) {
		m := r.(*Name)
		m.PolicyId = "pid_dbhid:xxxxx"
		m.Name = "new_name"
	})
}
