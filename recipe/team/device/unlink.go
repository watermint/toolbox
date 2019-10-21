package device

import (
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type UnlinkVO struct {
	DeleteOnUnlink bool
	File           app_file.Data
	Peer           app_conn.ConnBusinessFile
}

type UnlinkWorker struct {
	session *mo_device.MemberSession
	rep     app_report.Report
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

	err := sv_device.New(z.ctx).Revoke(z.session.Session())
	if err != nil {
		z.rep.Failure(app_msg.M("recipe.team.device.unlink.err.unable_to_unlink", app_msg.P{
			"Error": err.Error(),
		}), z.session, nil)
		return err
	}
	z.rep.Success(z.session, nil)
	return nil
}

type Unlink struct {
}

func (z *Unlink) Console() {
}

func (z *Unlink) Requirement() app_vo.ValueObject {
	return &UnlinkVO{}
}

func (z *Unlink) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UnlinkVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	err = vo.File.Model(k.Control(), &mo_device.MemberSession{})
	if err != nil {
		return err
	}

	rep, err := k.Report("unlink", app_report.TransactionHeader(&mo_device.MemberSession{}, nil))
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	err = vo.File.EachRow(func(m interface{}, rowIndex int) error {
		q.Enqueue(&UnlinkWorker{
			session: m.(*mo_device.MemberSession),
			rep:     rep,
			ctx:     ctx,
			ctl:     k.Control(),
		})
		return nil
	})
	q.Wait()
	return nil
}

func (z *Unlink) Test(c app_control.Control) error {
	return nil
}
