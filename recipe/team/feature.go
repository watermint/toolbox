package team

import (
	"github.com/watermint/toolbox/domain/model/mo_team"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type FeatureVO struct {
	PeerName app_conn.ConnBusinessInfo
}

func (*FeatureVO) Validate(t app_vo.Validator) {
}

type Feature struct {
}

func (*Feature) Requirement() app_vo.ValueObject {
	return &FeatureVO{}
}

func (*Feature) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*FeatureVO)
	conn, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	// Write report
	rep, err := k.Report("feature", &mo_team.Feature{})
	if err != nil {
		return err
	}
	defer rep.Close()

	info, err := sv_team.New(conn).Feature()
	if err != nil {
		return err
	}
	rep.Row(info)

	return nil
}
