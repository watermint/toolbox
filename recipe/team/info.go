package team

import (
	"github.com/watermint/toolbox/domain/model/mo_team"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type Info struct {
}

type InfoVO struct {
	PeerName app_conn.ConnBusinessInfo
}

func (*InfoVO) Validate(t app_vo.Validator) {
}

func (Info) Requirement() app_vo.ValueObject {
	return &InfoVO{}
}

func (Info) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*InfoVO)
	conn, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	// Write report
	rep, err := k.Report("info", &mo_team.Info{})
	if err != nil {
		return err
	}
	defer rep.Close()

	info, err := sv_team.New(conn).Info()
	if err != nil {
		return err
	}
	rep.Row(info)

	return nil
}
