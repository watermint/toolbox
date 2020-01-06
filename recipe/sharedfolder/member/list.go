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
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type ListWorker struct {
	folder *mo_sharedfolder.SharedFolder
	conn   api_context.Context
	rep    rp_model.RowReport
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	z.ctl.UI().InfoK("recipe.sharedfolder.member.list.progress.scan",
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
	Peer   rc_conn.ConnUserFile
	Member rp_model.RowReport
}

func (z *List) Preset() {
	z.Member.SetModel(&mo_sharedfolder_member.SharedFolderMember{})
}

func (z *List) Exec(c app_control.Control) error {
	folders, err := sv_sharedfolder.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.Member.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	for _, folder := range folders {
		q.Enqueue(&ListWorker{
			folder: folder,
			conn:   z.Peer.Context(),
			rep:    z.Member,
			ctl:    c,
		})
	}
	q.Wait()

	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "member", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
		}
		return nil
	})
}
