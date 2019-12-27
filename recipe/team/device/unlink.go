package device

import (
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type UnlinkVO struct {
}

type UnlinkWorker struct {
	session *mo_device.MemberSession
	rep     rp_model.TransactionReport
	ctx     api_context.Context
	ctl     app_control.Control
}

func (z *UnlinkWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Info("recipe.team.device.unlink.progress", app_msg.P{
		"Member":      z.session.Email,
		"SessionType": z.session.DeviceTag,
		"SessionId":   z.session.Id,
	})

	s := &mo_device.Metadata{
		Tag:          z.session.DeviceTag,
		TeamMemberId: z.session.TeamMemberId,
		Id:           z.session.Id,
	}
	err := sv_device.New(z.ctx).Revoke(s)
	if err != nil {
		z.rep.Failure(err, z.session)
		return err
	}
	z.rep.Success(z.session, nil)
	return nil
}

const (
	reportUnlink = "unlink"
)

type Unlink struct {
	DeleteOnUnlink bool
	File           fd_file.RowFeed
	Peer           rc_conn.ConnBusinessFile
	OperationLog   rp_model.TransactionReport
}

func (z *Unlink) Preset() {
	z.File.SetModel(&mo_device.MemberSession{})
	z.OperationLog.SetModel(&mo_device.MemberSession{}, nil)
}

func (z *Unlink) Console() {
}

func (z *Unlink) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	err := z.File.EachRow(func(m interface{}, rowIndex int) error {
		q.Enqueue(&UnlinkWorker{
			session: m.(*mo_device.MemberSession),
			rep:     z.OperationLog,
			ctx:     z.Peer.Context(),
			ctl:     c,
		})
		return nil
	})
	q.Wait()
	return err
}

func (z *Unlink) Test(c app_control.Control) error {
	return qt_endtoend.HumanInteractionRequired()
}
