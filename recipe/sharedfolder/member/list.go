package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
)

type ListVO struct {
	Peer app_conn.ConnUserFile
}

type ListWorker struct {
	folder *mo_sharedfolder.SharedFolder
	conn   api_context.Context
	rep    app_report.Report
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	z.ctl.UI().Info("recipe.sharedfolder.member.list.progress.scan",
		app_msg.P{
			"Folder":   z.folder.Name,
			"FolderId": z.folder.SharedFolderId},
	)
	z.ctl.Log().Debug("Scanning folder", zap.Any("folder", z.folder))
	members, err := sv_sharedfolder_member.New(z.conn, z.folder).List()
	if err != nil {
		return err
	}

	for _, member := range members {
		z.rep.Row(mo_sharedfolder_member.NewSharedFolderMember(z.folder, member))
	}
	return nil
}

type List struct {
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (*List) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*ListVO)
	conn, err := lvo.Peer.Connect(k.Control())
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

	q := k.NewQueue()
	for _, folder := range folders {
		q.Enqueue(&ListWorker{
			folder: folder,
			conn:   conn,
			rep:    rep,
			ctl:    k.Control(),
		})
	}
	q.Wait()

	return nil
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
