package member

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Newlist struct {
	Peer       rc_conn.ConnBusinessInfo
	MemberList rp_model.RowReport
}

func (z *Newlist) Hidden() {
}

func (z *Newlist) Exec(k rc_kitchen.Kitchen) error {
	ctx, err := z.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	if err = z.MemberList.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	for _, member := range members {
		z.MemberList.Row(member)
	}
	return nil
}

func (z *Newlist) Test(c app_control.Control) error {
	z.Init()
	return z.Exec(rc_kitchen.NewKitchen(c, z))
}

func (z *Newlist) Init() {
	z.MemberList.Model(&mo_member.Member{})
}
