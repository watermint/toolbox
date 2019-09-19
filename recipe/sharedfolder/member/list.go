package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"go.uber.org/zap"
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
	return app_test.TestRows(c, "sharedfolder_member", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
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

	folders, err := sv_sharedfolder.New(conn).List()
	if err != nil {
		return err
	}

	rep, err := k.Report("sharedfolder_member", &mo_sharedfolder_member.SharedFolderMember{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, folder := range folders {
		//k.UI().Info("progress.scan",
		//	app_msg.P("Folder", folder.Name),
		//	app_msg.P("FolderId", folder.SharedFolderId),
		//)
		k.Log().Debug("Scanning folder", zap.Any("folder", folder))
		members, err := sv_sharedfolder_member.New(conn, folder).List()
		if err != nil {
			return err
		}

		for _, member := range members {
			rep.Row(mo_sharedfolder_member.NewSharedFolderMember(folder, member))
		}
	}
	return nil
}
