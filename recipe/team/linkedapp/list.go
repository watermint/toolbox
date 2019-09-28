package linkedapp

import (
	"github.com/watermint/toolbox/domain/model/mo_linkedapp"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_linkedapp"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type ListVO struct {
	Peer app_conn.ConnBusinessFile
}

type List struct {
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	lvo := k.Value().(*ListVO)
	conn, err := lvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	memberList, err := sv_member.New(conn).List()
	if err != nil {
		return err
	}
	members := mo_member.MapByTeamMemberId(memberList)

	rep, err := k.Report("linkedapp", &mo_linkedapp.MemberLinkedApp{})
	if err != nil {
		return err
	}
	defer rep.Close()

	apps, err := sv_linkedapp.New(conn).List()
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

		rep.Row(ma)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return nil
}
