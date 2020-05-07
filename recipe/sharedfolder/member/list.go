package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type MsgList struct {
	ProgressScan app_msg.Message
}

var (
	MList = app_msg.Apply(&MsgList{}).(*MsgList)
)

type ListWorker struct {
	folder *mo_sharedfolder.SharedFolder
	conn   dbx_context.Context
	rep    rp_model.RowReport
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	z.ctl.UI().Progress(MList.ProgressScan.With("Folder", z.folder.Name).With("FolderId", z.folder.SharedFolderId))

	z.ctl.Log().Debug("Scanning folder", es_log.Any("folder", z.folder))
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
	Peer   dbx_conn.ConnUserFile
	Member rp_model.RowReport
}

func (z *List) Preset() {
	z.Member.SetModel(
		&mo_sharedfolder_member.SharedFolderMember{},
		rp_model.HiddenColumns(
			"shared_folder_id",
			"parent_shared_folder_id",
			"account_id",
			"group_id",
		),
	)
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
	return qtr_endtoend.TestRows(c, "member", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
		}
		return nil
	})
}
