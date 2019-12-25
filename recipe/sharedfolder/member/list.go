package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type ListVO struct {
	Peer rc_conn.OldConnUserFile
}

type ListWorker struct {
	folder *mo_sharedfolder.SharedFolder
	conn   api_context.Context
	rep    rp_model.SideCarReport
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

const (
	reportList = "sharedfolder_member"
)

type List struct {
}

func (z *List) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportList, &mo_sharedfolder_member.SharedFolderMember{}),
	}
}

func (z *List) Requirement() rc_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
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

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportList)
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
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "sharedfolder_member", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
		}
		return nil
	})
}
