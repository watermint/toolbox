package dev

import (
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/recpie/app_worker_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type AsyncVO struct {
	PeerName app_conn.ConnBusinessMgmt
}

type Async struct {
}

func (z *Async) Hidden() {
}

func (z *Async) Requirement() app_vo.ValueObject {
	return &AsyncVO{}
}

func (z *Async) Exec(k app_kitchen.Kitchen) error {
	ui := k.UI()
	var vo interface{} = k.Value()
	lvo := vo.(*AsyncVO)
	connInfo, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	q := app_worker_impl.NewQueue(k.Control())
	q.Launch(4)

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
		q.Enqueue(func(ctl app_control.Control) error {
			ui.Info("recipe.group.member.list.progress.scan", app_msg.P{"Group": group.GroupName})

			msv := sv_group_member.New(connInfo, group)
			members, err := msv.List()
			if err != nil {
				return err
			}
			for _, m := range members {
				row := mo_group_member.NewGroupMember(group, m)
				rep.Row(row)
			}
			return nil
		})
	}
	q.Wait()

	return nil
}

func (z *Async) Test(c app_control.Control) error {
	return z.Exec(app_kitchen.NewKitchen(c, &app_vo.EmptyValueObject{}))
}
