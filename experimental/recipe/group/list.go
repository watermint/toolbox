package group

import (
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/experimental/app_conn"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_vo"
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

	groups, err := sv_group.New(connInfo).List()
	if err != nil {
		return err
	}
	rep, err := k.Report("group", &mo_group.Group{})
	if err != nil {
		return err
	}
	defer rep.Close()
	for _, m := range groups {
		rep.Row(m)
	}
	return nil
}
