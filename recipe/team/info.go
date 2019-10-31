package team

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_team"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type Info struct {
}

func (z *Info) Test(c app_control.Control) error {
	lvo := &InfoVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return qt_test.NotEnoughResource()
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "info", func(cols map[string]string) error {
		if _, ok := cols["team_id"]; !ok {
			return errors.New("`team_id` is not found")
		}
		return nil
	})
}

type InfoVO struct {
	Peer app_conn.ConnBusinessInfo
}

func (Info) Requirement() app_vo.ValueObject {
	return &InfoVO{}
}

func (Info) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*InfoVO)
	conn, err := lvo.Peer.Connect(k.Control())
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
