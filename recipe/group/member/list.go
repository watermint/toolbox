package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type ListVO struct {
	PeerName app_conn.ConnBusinessInfo
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
	return app_test.TestRows(c, "group_member", func(cols map[string]string) error {
		if _, ok := cols["group_id"]; !ok {
			return errors.New("group_id is not found")
		}
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("team_member_id is not found")
		}
		return nil
	})
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	ui := k.UI()
	var vo interface{} = k.Value()
	lvo := vo.(*ListVO)
	connInfo, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	gsv := sv_group.New(connInfo)
	groups, err := gsv.List()
	if err != nil {
		return err
	}

	rep, err := k.Report("group_member", &mo_group_member.GroupMember{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, group := range groups {
		ui.Info("recipe.group.member.progress.scan", app_msg.P{"Group": group.GroupName})

		msv := sv_group_member.New(connInfo, group)
		members, err := msv.List()
		if err != nil {
			return err
		}
		for _, m := range members {
			row := mo_group_member.NewGroupMember(group, m)
			rep.Row(row)
		}
	}
	return nil
}
