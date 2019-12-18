package device

import (
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type UnlinkVO struct {
	DeleteOnUnlink bool
	File           fd_file.Feed
	Peer           rc_conn.ConnBusinessFile
}

type UnlinkWorker struct {
	session *mo_device.MemberSession
	rep     rp_model.Report
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
}

func (z *Unlink) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportUnlink,
			rp_model.TransactionHeader(&mo_device.MemberSession{}, nil)),
	}
}

func (z *Unlink) Console() {
}

func (z *Unlink) Requirement() rc_vo.ValueObject {
	return &UnlinkVO{}
}

func (z *Unlink) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*UnlinkVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	err = vo.File.Model(k.Control(), &mo_device.MemberSession{})
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportUnlink)
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
	return qt_recipe.HumanInteractionRequired()
}
