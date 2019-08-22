package member

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type List struct {
}

type ListVO struct {
	PeerName app_conn.ConnBusinessInfo
}

func (*ListVO) Validate(t app_vo.Validator) {
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

	members, err := sv_member.New(connInfo).List()
	if err != nil {
		return err
	}

	rep, err := k.Report("member", &mo_member.Member{})
	if err != nil {
		return err
	}
	defer rep.Close()
	for _, m := range members {
		rep.Row(m)
	}
	return nil
}
