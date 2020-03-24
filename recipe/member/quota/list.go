package quota

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member_quota"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member_quota"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
)

type ListWorker struct {
	member *mo_member.Member
	ctx    api_context.Context
	rep    rp_model.RowReport
	ctl    app_control.Control
}

func (z *ListWorker) Exec() error {
	l := z.ctl.Log()

	z.ctl.UI().InfoK("recipe.member.quota.list.scan", app_msg.P{"MemberEmail": z.member.Email})
	l.Debug("Scan member", zap.String("Routine", ut_runtime.GetGoRoutineName()), zap.Any("Member", z.member))

	q, err := sv_member_quota.NewQuota(z.ctx).Resolve(z.member.TeamMemberId)
	if err != nil {
		return err
	}
	z.rep.Row(mo_member_quota.NewMemberQuota(z.member, q))
	return nil
}

type List struct {
	Peer        rc_conn.ConnBusinessMgmt
	MemberQuota rp_model.RowReport
}

func (z *List) Preset() {
	z.MemberQuota.SetModel(&mo_member_quota.MemberQuota{})
}

func (z *List) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	if err := z.MemberQuota.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	for _, member := range members {
		q.Enqueue(&ListWorker{
			member: member,
			ctx:    z.Peer.Context(),
			rep:    z.MemberQuota,
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
	return qt_recipe.TestRows(c, "member_quota", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("`email` is not found")
		}
		if _, ok := cols["quota"]; !ok {
			return errors.New("`quota` is not found")
		}
		return nil
	})
}
