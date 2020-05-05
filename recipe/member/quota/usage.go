package quota

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_usage"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_usage"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type MsgUsage struct {
	ProgressScan app_msg.Message
}

var (
	MUsage = app_msg.Apply(&MsgUsage{}).(*MsgUsage)
)

type UsageVO struct {
}

type UsageWorker struct {
	member *mo_member.Member
	ctx    dbx_context.Context
	ctl    app_control.Control
	rep    rp_model.RowReport
}

func (z *UsageWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Progress(MUsage.ProgressScan.With("MemberEmail", z.member.Email))
	l := z.ctl.Log().With(es_log.Any("member", z.member))
	l.Debug("Scanning")

	usage, err := sv_usage.New(z.ctx.AsMemberId(z.member.TeamMemberId)).Resolve()
	if err != nil {
		l.Debug("Unable to scan usage data", es_log.Error(err))
		return err
	}

	z.rep.Row(mo_usage.NewMemberUsage(z.member, usage))
	return nil
}

type Usage struct {
	Peer  dbx_conn.ConnBusinessFile
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
