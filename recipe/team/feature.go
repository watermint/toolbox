package team

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_team"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type FeatureVO struct {
	Peer app_conn.ConnBusinessInfo
}

type Feature struct {
}

func (z *Feature) Test(c app_control.Control) error {
	lvo := &FeatureVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "feature", func(cols map[string]string) error {
		if _, ok := cols["upload_api_rate_limit"]; !ok {
			return errors.New("`upload_api_rate_limit` is not found")
		}
		return nil
	})
}

func (*Feature) Requirement() app_vo.ValueObject {
	return &FeatureVO{}
}

func (*Feature) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*FeatureVO)
	conn, err := lvo.Peer.Connect(k.Control())
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
