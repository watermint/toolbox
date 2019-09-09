package sharedfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type ListVO struct {
	PeerName app_conn.ConnUserFile
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
	return app_test.TestRows(c, "sharedfolder", func(cols map[string]string) error {
		if _, ok := cols["SharedFolderId"]; !ok {
			return errors.New("SharedFolderId is not found")
		}
		return nil
	})
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (*List) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*ListVO)
	conn, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	k.Log().Debug("Scanning folders")
	folders, err := sv_sharedfolder.New(conn).List()
	if err != nil {
		return err
	}

	// Write report
	rep, err := k.Report("sharedfolder", &mo_sharedfolder.SharedFolder{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, folder := range folders {
		rep.Row(folder)
	}
	return nil
}
