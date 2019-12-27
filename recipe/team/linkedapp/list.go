package linkedapp

import (
	"github.com/watermint/toolbox/domain/model/mo_linkedapp"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_linkedapp"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer      rc_conn.ConnBusinessFile
	LinkedApp rp_model.RowReport
}

func (z *List) Preset() {
	z.LinkedApp.SetModel(&mo_linkedapp.MemberLinkedApp{})
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
	memberList, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	members := mo_member.MapByTeamMemberId(memberList)

	if err := z.LinkedApp.Open(); err != nil {
		return err
	}

	apps, err := sv_linkedapp.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	for _, app := range apps {
		m := &mo_member.Member{}
		m.TeamMemberId = app.TeamMemberId

		if m0, e := members[app.TeamMemberId]; e {
			m = m0
		}
		ma := mo_linkedapp.NewMemberLinkedApp(m, app)

		z.LinkedApp.Row(ma)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return nil
}
