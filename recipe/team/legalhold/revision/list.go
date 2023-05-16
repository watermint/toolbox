package revision

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_legalhold"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_legalhold"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"time"
)

type List struct {
	Peer     dbx_conn.ConnScopedTeam
	PolicyId string
	Revision rp_model.RowReport
	After    mo_time.Time
}

func (z *List) Preset() {
	//z.Peer.SetScopes(
	//dbx_auth.ScopeMembersRead,
	//dbx_auth.ScopeTeamDataGovernanceWrite,
	//)
	z.Revision.SetModel(&mo_legalhold.Revision{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Revision.Open(); err != nil {
		return err
	}
	err := sv_legalhold.New(z.Peer.Client()).Revisions(z.PolicyId, z.After.Time(), func(rev *mo_legalhold.Revision) {
		z.Revision.Row(rev)
	})
	return err
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.PolicyId = "pid_dbhid:xxxxx"
		m.After = mo_time.New(time.Now().Add(-7 * time.Hour * 24))
	})
}
