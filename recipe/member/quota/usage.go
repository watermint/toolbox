package quota

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_usage"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_usage"
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

type UsageVO struct {
	Peer rc_conn.ConnBusinessFile
}

type UsageWorker struct {
	member *mo_member.Member
	ctx    api_context.Context
	ctl    app_control.Control
	rep    rp_model.Report
}

func (z *UsageWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Info("recipe.member.quota.usage.scan",
		app_msg.P{
			"MemberEmail": z.member.Email,
		})
	l := z.ctl.Log().With(zap.Any("member", z.member))
	l.Debug("Scanning")

	usage, err := sv_usage.New(z.ctx.AsMemberId(z.member.TeamMemberId)).Resolve()
	if err != nil {
		l.Debug("Unable to scan usage data", zap.Error(err))
		return err
	}

	z.rep.Row(mo_usage.NewMemberUsage(z.member, usage))
	return nil
}

const (
	reportUsage = "usage"
)

type Usage struct {
}

func (z *Usage) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportUsage, &mo_usage.MemberUsage{}),
	}
}

func (z *Usage) Requirement() rc_vo.ValueObject {
	return &UsageVO{}
}

func (z *Usage) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*UsageVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportUsage)
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	for _, member := range members {
		q.Enqueue(&UsageWorker{
			member: member,
			ctx:    ctx,
			ctl:    k.Control(),
			rep:    rep,
		})
	}
	q.Wait()
	return nil
}

func (z *Usage) Test(c app_control.Control) error {
	vo := &UsageVO{}
	if !qt_recipe.ApplyTestPeers(c, vo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, vo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "usage", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		if _, ok := cols["used_bytes"]; !ok {
			return errors.New("`used_bytes` is not found")
		}
		return nil
	})
}
