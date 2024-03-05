package legalhold

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_legalhold"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_legalhold"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer            dbx_conn.ConnScopedTeam
	Policies        rp_model.RowReport
	IncludeReleased bool
}

func (z *List) Preset() {
	//z.Peer.SetScopes(
	//	dbx_auth.ScopeTeamDataGovernanceWrite,
	//)
	z.Policies.SetModel(&mo_legalhold.Policy{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Policies.Open(); err != nil {
		return err
	}
	policies, err := sv_legalhold.New(z.Peer.Client()).List(z.IncludeReleased)
	if err != nil {
		return err
	}
	for _, policy := range policies {
		z.Policies.Row(policy)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
