package quota

import (
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_member_quota"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_member_quota"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
)

type UpdateVO struct {
	PeerName app_conn.ConnBusinessMgmt
	File     app_file.Data
}

type UpdateWorker struct {
	member *mo_member.Member
	quota  int

	ctl app_control.Control
	ctx api_context.Context
	rep app_report.Report
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
		z.rep.Failure(
			app_msg.M("recipe.member.quota.update.err.cannot_update", app_msg.P{"Error": err.Error()}),
			in,
			nil,
		)
	} else {
		z.rep.Success(in, mo_member_quota.NewMemberQuota(z.member, newQuota))
	}
	return nil
}

type Update struct {
}

func (z *Update) Console() {
}

func (z *Update) Requirement() app_vo.ValueObject {
	return &UpdateVO{}
}

func (z *Update) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UpdateVO)

	ctx, err := vo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	rep, err := k.Report("quota_update",
		app_report.TransactionHeader(
			&mo_member_quota.MemberQuota{},
			&mo_member_quota.MemberQuota{},
		),
	)
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
			rep.Failure(
				app_msg.M("recipe.member.quota.update.err.member_not_found_for_email",
					app_msg.P{"Email": mq.Email}),
				mq,
				nil,
			)
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
	return nil
}
