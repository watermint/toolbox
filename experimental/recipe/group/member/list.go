package member

import (
	"github.com/watermint/toolbox/domain/model/mo_group_member"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/experimental/app_conn"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_msg"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type ListVO struct {
	PeerName app_conn.ConnBusinessInfo
}

func (*ListVO) Validate(t app_vo.Validator) {
}

type List struct {
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (*List) Exec(k app_kitchen.Kitchen) error {
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
		k.UI().Info("recipe.group.member.list.progress.scan", app_msg.P("Group", group.GroupName))

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
