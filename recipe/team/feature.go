package team

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_team"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type FeatureVO struct {
	Peer rc_conn.ConnBusinessInfo
}

const (
	reportFeature = "feature"
)

type Feature struct {
}

func (z *Feature) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportFeature, &mo_team.Feature{}),
	}
}

func (z *Feature) Test(c app_control.Control) error {
	lvo := &FeatureVO{}
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "feature", func(cols map[string]string) error {
		if _, ok := cols["upload_api_rate_limit"]; !ok {
			return errors.New("`upload_api_rate_limit` is not found")
		}
		return nil
	})
}

func (z *Feature) Requirement() rc_vo.ValueObject {
	return &FeatureVO{}
}

func (z *Feature) Exec(k rc_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*FeatureVO)
	conn, err := lvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	// Write report
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportFeature)
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
