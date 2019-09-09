package sharedlink

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
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

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "sharedlink", func(cols map[string]string) error {
		if _, ok := cols["SharedLinkId"]; !ok {
			return errors.New("`SharedLinkId` is not found")
		}
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("`team_member_id` is not found")
		}
		return nil
	})
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (*List) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*ListVO)
	conn, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(conn).List()
	if err != nil {
		return err
	}

	// Write report
	rep, err := k.Report("sharedlink", &mo_sharedlink.SharedLinkMember{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, member := range members {
		mc := conn.AsMemberId(member.TeamMemberId)
		links, err := sv_sharedlink.New(mc).List()
		if err != nil {
			return err
		}
		for _, link := range links {
			lm := mo_sharedlink.NewSharedLinkMember(link, member)
			rep.Row(lm)
		}
	}
	return nil
}