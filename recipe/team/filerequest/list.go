package filerequest

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
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
	member *mo_member.Member
	conn   dbx_context.Context
	rep    rp_model.RowReport
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	z.ctl.UI().Progress(MList.ProgressScan.With("MemberEmail", z.member.Email))
	mc := z.conn.AsMemberId(z.member.TeamMemberId)
	reqs, err := sv_filerequest.New(mc).List()
	if err != nil {
		return err
	}
	for _, req := range reqs {
		fm := mo_filerequest.NewMemberFileRequest(req, z.member)
		z.rep.Row(fm)
	}
	return nil
}

type List struct {
	Peer        dbx_conn.ConnBusinessFile
	FileRequest rp_model.RowReport
}

func (z *List) Preset() {
	z.FileRequest.SetModel(
		&mo_filerequest.MemberFileRequest{},
		rp_model.HiddenColumns(
			"account_id",
			"team_member_id",
			"file_request_id",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.FileRequest.Open(); err != nil {
		return err
	}

	q := c.NewLegacyQueue()
	for _, member := range members {
		q.Enqueue(&ListWorker{
			member: member,
			conn:   z.Peer.Context(),
			rep:    z.FileRequest,
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
	return qtr_endtoend.TestRows(c, "file_request", func(cols map[string]string) error {
		if _, ok := cols["url"]; !ok {
			return errors.New("`url` is not found")
		}
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		return nil
	})
}
