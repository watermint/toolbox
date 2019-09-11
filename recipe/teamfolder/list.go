package teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"go.uber.org/zap"
)

type ListVO struct {
	PeerName app_conn.ConnBusinessFile
}

type List struct {
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "teamfolder", func(cols map[string]string) error {
		if _, ok := cols["team_folder_id"]; !ok {
			return errors.New("`team_folder_id` is not found")
		}
		return nil
	})
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	// TypeAssertionError will be handled by infra
	var vo interface{} = k.Value()
	fvo := vo.(*ListVO)

	connFile, err := fvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	folders, err := sv_teamfolder.New(connFile).List()
	if err != nil {
		// ApiError will be reported by infra
		return err
	}

	rep, err := k.Report("teamfolder", &mo_teamfolder.TeamFolder{})
	if err != nil {
		return err
	}
	defer rep.Close()
	for _, folder := range folders {
		k.Log().Debug("Folder", zap.Any("folder", folder))
		rep.Row(folder)
	}

	return nil
}
