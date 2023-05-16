package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_legalhold"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer     dbx_conn.ConnScopedTeam
	PolicyId string
	Member   rp_model.RowReport
}

func (z *List) Preset() {
	//z.Peer.SetScopes(
	//dbx_auth.ScopeMembersRead,
	//dbx_auth.ScopeTeamDataGovernanceWrite,
	//)
	z.Member.SetModel(&mo_member.Member{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Member.Open(); err != nil {
		return err
	}
	policy, err := sv_legalhold.New(z.Peer.Client()).Info(z.PolicyId)
	if err != nil {
		return err
	}
	svm := sv_member.NewCached(z.Peer.Client())
	for _, memberId := range policy.TeamMemberIds() {
		member, err := svm.Resolve(memberId)
		if err != nil {
			return err
		}
		z.Member.Row(member)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.PolicyId = "pid_dbhid:xxxxx"
	})
}
