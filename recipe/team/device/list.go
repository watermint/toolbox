package device

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type ListVO struct {
	PeerName app_conn.ConnBusinessFile
}

type List struct {
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	lvo := k.Value().(*ListVO)
	ctx, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	memberList, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	members := mo_member.MapByTeamMemberId(memberList)

	sessions, err := sv_device.New(ctx).List()
	if err != nil {
		return err
	}

	rep, err := k.Report("device", &mo_device.MemberSession{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, session := range sessions {
		if m, e := members[session.EntryTeamMemberId()]; e {
			ma := mo_device.NewMemberSession(m, session)
			rep.Row(ma)
		}
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
	return app_test.TestRows(c, "device", func(cols map[string]string) error {
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("team_member_id is not found")
		}
		return nil
	})
}
