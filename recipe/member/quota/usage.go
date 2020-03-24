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
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type UsageVO struct {
}

type UsageWorker struct {
	member *mo_member.Member
	ctx    api_context.DropboxApiContext
	ctl    app_control.Control
	rep    rp_model.RowReport
}

func (z *UsageWorker) Exec() error {
	ui := z.ctl.UI()
	ui.InfoK("recipe.member.quota.usage.scan",
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

type Usage struct {
	Peer  rc_conn.ConnBusinessFile
	Usage rp_model.RowReport
}

func (z *Usage) Preset() {
	z.Usage.SetModel(&mo_usage.MemberUsage{})
}

func (z *Usage) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.Usage.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	for _, member := range members {
		q.Enqueue(&UsageWorker{
			member: member,
			ctx:    z.Peer.Context(),
			ctl:    c,
			rep:    z.Usage,
		})
	}
	q.Wait()
	return nil
}

func (z *Usage) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &Usage{}, rc_recipe.NoCustomValues); err != nil {
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
