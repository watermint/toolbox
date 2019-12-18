package quota

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_member_quota"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_member_quota"
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
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type UpdateVO struct {
	Peer rc_conn.ConnBusinessMgmt
	File fd_file.Feed
}

type UpdateWorker struct {
	member *mo_member.Member
	quota  int

	ctl app_control.Control
	ctx api_context.Context
	rep rp_model.Report
}

func (z *UpdateWorker) Exec() error {
	l := z.ctl.Log()
	z.ctl.UI().Info("recipe.member.quota.update.progress",
		app_msg.P{
			"MemberEmail": z.member.Email,
			"Quota":       z.quota,
		})
	l.Debug("Updating quota", zap.String("Routine", ut_runtime.GetGoRoutineName()), zap.Any("Member", z.member))

	q := &mo_member_quota.Quota{
		TeamMemberId: z.member.TeamMemberId,
		Quota:        z.quota,
	}
	in := mo_member_quota.MemberQuota{
		Email: z.member.Email,
		Quota: z.quota,
	}

	newQuota, err := sv_member_quota.NewQuota(z.ctx).Update(q)
	if err != nil {
		z.rep.Failure(err, in)
	} else {
		z.rep.Success(in, mo_member_quota.NewMemberQuota(z.member, newQuota))
	}
	return nil
}

const (
	reportUpdate = "quota_update"
)

type Update struct {
}

func (z *Update) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportUpdate, rp_model.TransactionHeader(
			&mo_member_quota.MemberQuota{},
			&mo_member_quota.MemberQuota{},
		)),
	}
}

func (z *Update) Console() {
}

func (z *Update) Requirement() rc_vo.ValueObject {
	return &UpdateVO{}
}

func (z *Update) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*UpdateVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportUpdate)
	if err != nil {
		return err
	}
	defer rep.Close()

	if err := vo.File.Model(k.Control(), &mo_member_quota.MemberQuota{}); err != nil {
		return err
	}

	q := k.NewQueue()

	vo.File.EachRow(func(m interface{}, rowIndex int) error {
		mq := m.(*mo_member_quota.MemberQuota)
		member, ok := emailToMember[mq.Email]
		if !ok {
			rep.Failure(&rp_model.NotFound{Id: mq.Email}, mq)
			return nil
		}

		q.Enqueue(&UpdateWorker{
			member: member,
			quota:  mq.Quota,
			ctl:    k.Control(),
			ctx:    ctx,
			rep:    rep,
		})
		return nil
	})
	q.Wait()

	return nil
}

func (z *Update) Test(c app_control.Control) error {
	return qt_recipe.HumanInteractionRequired()
}
