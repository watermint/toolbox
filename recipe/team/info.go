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

const (
	reportInfo = "info"
)

type Info struct {
}

func (z *Info) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportInfo, &mo_team.Info{}),
	}
}

func (z *Info) Test(c app_control.Control) error {
	lvo := &InfoVO{}
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "info", func(cols map[string]string) error {
		if _, ok := cols["team_id"]; !ok {
			return errors.New("`team_id` is not found")
		}
		return nil
	})
}

type InfoVO struct {
	Peer rc_conn.ConnBusinessInfo
}

func (z *Info) Requirement() rc_vo.ValueObject {
	return &InfoVO{}
}

func (z *Info) Exec(k rc_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*InfoVO)
	conn, err := lvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	// Write report
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportInfo)
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
